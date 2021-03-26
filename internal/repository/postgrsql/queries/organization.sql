/*
 ORGANIZATIONS #########################################################################################################
 */

/*
 Удалить / Архивировать - означает что у записи поле entry_deleted_date_time не равно null и выставлено время удаления
*/

/*
 * Данный sql запрос создает таблицу 'organizations'.
 * В таблице хранится информация касательно всех компаний 'organizations'.
 */
create table organizations (
    entry_created_date_time timestamp not null default now(),
    entry_updated_date_time timestamp not null default now(),
    entry_deleted_date_time timestamp null,
    organization_id uuid not null default uuid_generate_v4() primary key,
    organization_name varchar not null,
    organization_description text null,
    parent_organization_id uuid null,
    foreign key (parent_organization_id) references organizations (organization_id),
    parent_organization_name varchar null,
    parent_organization_description text null,
    root_organization_id uuid null,
    foreign key (root_organization_id) references organizations (organization_id),
    root_organization_name varchar null,
    root_organization_description text null
);

/*
 * Данный sql запрос создает триггер на операцию INSERT используя функцию созданную выше update_organization()
 * для обновления записи в таблице 'organizations'.
 */
create trigger organization_updater
after insert
on public.organizations
for each row
execute procedure update_organization();

/*
 * Данный sql запрос создает функцию для обновления записи в таблице 'organizations'.
 * Автоматический устанавливаются значения parent_organization_name, root_organization_id,
 * root_organization_name, organization_level, organization_level, root_organization_level,
 * tree_organization_id и tree_organization_name на основе organization_name и parent_organization_id.
 */
create or replace function update_organization()
returns trigger as
$BODY$
begin
    update organizations set
        parent_organization_name = (select organization_name from organizations where organization_id = new.parent_organization_id),
        root_organization_id = (select root_organization_id from organizations where organization_id = new.parent_organization_id),
        root_organization_name = (select root_organization_name from organizations where organization_id = new.parent_organization_id),
        organization_level = ((select organization_level from organizations where organization_id = new.parent_organization_id) + 1),
        parent_organization_level = (select organization_level from organizations where organization_id = new.parent_organization_id),
        root_organization_level = 1,
        tree_organization_id = ((select tree_organization_id from organizations where organization_id = new.parent_organization_id) || '\' || new.organization_id::text),
        tree_organization_name = ((select tree_organization_name from organizations where organization_id = new.parent_organization_id) || '\' || new.organization_name::text)
    where organization_id = new.organization_id;
    return new;
end;
$BODY$
language 'plpgsql';
