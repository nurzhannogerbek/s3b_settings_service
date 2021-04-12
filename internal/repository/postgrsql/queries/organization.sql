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
    if new.parent_organization_id is null then
        update organizations set
             parent_organization_name = new.organization_name,
             parent_organization_id = new.organization_id,
             root_organization_id = new.organization_id,
             root_organization_name = new.organization_name,
             organization_level = 1,
             parent_organization_level = 1,
             root_organization_level = 1,
             tree_organization_id = '\' || new.organization_id::text,
             tree_organization_name = '\' || new.organization_name::text
        where organization_id = new.organization_id;
    else
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
    end if;
    return new;
end;
$BODY$
language 'plpgsql';

/*
 * Запрос на удаление функции описанной выше
 */
drop function public.update_organization();

/*
 * Запрос на создание триггера на функцию update_organization
 */
create trigger organization_updater
after insert
on public.organizations
for each row
execute procedure update_organization();

/*
 * Запрос на удаление триггера описанной выше
 */
drop trigger organization_updater on public.organizations;

/*
 * Данный запрос создает новую запись организации (organization), функция CreateOrganization
 */
insert into organizations (
    organization_name,
    parent_organization_id)
values (
    :organization_name,
    null)
returning organization_id,
    organization_name,
    parent_organization_id,
    parent_organization_name,
    root_organization_id,
    root_organization_name,
    organization_level,
    parent_organization_level,
    root_organization_level
    tree_organization_id,
    tree_organization_name;

/*
 * Данный запрос создает новый отдел организации (organization department), функция CreateOrganizationDepartment
 */
insert into organizations (
    organization_name,
    parent_organization_id)
values (
    :organization_name,
    :parent_organization_id)
returning organization_id,
    organization_name,
    parent_organization_id,
    parent_organization_name,
    root_organization_id,
    root_organization_name,
    organization_level,
    parent_organization_level,
    root_organization_level,
    tree_organization_id,
    tree_organization_name;

/*
 * Данный запрос архивирует организации (organization), функция DeleteOrganizations
 */
update
    organizations
set
    entry_deleted_date_time = now()
where
    organization_id in (?);

/*
 * Данный запрос запрашивает запись организации (organization) по id, функция GetOrganizationByID
 */
select
    organization_id,
    organization_name,
    parent_organization_id,
    parent_organization_name,
    root_organization_id,
    root_organization_name,
    organization_level,
    parent_organization_level,
    root_organization_level,
    tree_organization_id,
    tree_organization_name
from
    organizations
where
    organization_id = $1;

/*
 * Данный запрос запрашивает массив организаций (organizations) no ids, функция GetOrganizationsByIDs
 */
select
    organization_id,
    organization_name,
    parent_organization_id,
    parent_organization_name,
    root_organization_id,
    root_organization_name,
    organization_level,
    parent_organization_level,
    root_organization_level,
    tree_organization_id,
    tree_organization_name
from
    organizations
where
    organization_id in (?);

/*
 * Данный запрос обновляет название организации (organization) по id, функция UpdateOrganizationName
 */
update
    organizations
set entry_updated_date_time = now(),
    organization_name = $1
where organization_id = $2;

/*
 * Данный запрос обновляет поле tree_organization_name в организации no id, функция UpdateTreeOrganizationName
 */
update
    organizations
set entry_updated_date_time = now(),
    tree_organization_name = $1
where organization_id = $2;

/*
 * Данный запрос запрашивает список организаций имеющие в названии организации которое будет обновлено
 */
select organization_id,
       tree_organization_id
from organizations
where root_organization_id = (select root_organization_id
                              from organizations
                              where organization_id = $1::uuid)
    and tree_organization_name like '%' || '\' || $2::text || '%';

/*
 * Данный запрос восстанавливает организации по ids, функция RestoreDeletedOrganizations
 */
update organizations
set
    entry_deleted_date_time = null
where
    organization_id in ($1);

/*
 * Данный запрос запрашивает список отделов (department) принадлежащий определенной организации (organization)
 */
select
    organization_id,
    organization_name,
    parent_organization_id,
    parent_organization_name,
    root_organization_id,
    root_organization_name,
    organization_level,
    parent_organization_level,
    root_organization_level,
    tree_organization_id,
    tree_organization_name
from
    organizations
where
    parent_organization_id = $1
    and entry_deleted_date_time is null;

/*
 * Данный запрос запрашивает список всех отделов (department) принадлежащий основной организации (organization)
 */
select
    organization_id,
    organization_name,
    parent_organization_id,
    parent_organization_name,
    root_organization_id,
    root_organization_name,
    organization_level,
    parent_organization_level,
    root_organization_level,
    tree_organization_id,
    tree_organization_name
from
    organizations
where
    root_organization_id = $1
    and entry_deleted_date_time is null;

/*
 * Данный запрос запрашивает список архивированных отделов (department) принадлежащий определенной организации (organization)
 */
select
    organization_id,
    organization_name,
    parent_organization_id,
    parent_organization_name,
    root_organization_id,
    root_organization_name,
    organization_level,
    parent_organization_level,
    root_organization_level,
    tree_organization_id,
    tree_organization_name
from
    organizations
where
    parent_organization_id = $1
    and entry_deleted_date_time is not null;

/*
 * Данный запрос запрашивает список всех архивированных отделов (department) принадлежащий основной организации (organization)
 */
select
    organization_id,
    organization_name,
    parent_organization_id,
    parent_organization_name,
    root_organization_id,
    root_organization_name,
    organization_level,
    parent_organization_level,
    root_organization_level,
    tree_organization_id,
    tree_organization_name
from
    organizations
where
    root_organization_id = $1
    and entry_deleted_date_time is not null;