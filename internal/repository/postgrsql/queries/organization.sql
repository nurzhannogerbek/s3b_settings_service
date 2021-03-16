/*
 ORGANIZATIONS_SETTINGS #######################################################################################################
 */

/*
 Удалить / Архивировать - означает что у записи поле entry_deleted_date_time не равно null и выставлено время удаления
*/

/*
 * Данный sql запрос создает таблицу 'organizations_settings'.
 * В таблице хранится информация касательно настроек определенной компании 'organizations'.
 */
create table organizations_settings (
    entry_created_date_time timestamp not null default now(),
    entry_updated_date_time timestamp not null default now(),
    entry_deleted_date_time timestamp null,
    organization_id uuid null,
    foreign key (organization_id) references organizations (organization_id),
    country_id uuid null,
    foreign key (country_id) references countries (country_id),
    location_id uuid null,
    foreign key (location_id) references locations (location_id),
    organization_setting_address text null,
    organization_setting_postal_code varchar (50) null,
    organization_setting_work_time text null,
    organization_setting_privacy text not null,
    timezone_id uuid null,
    foreign key (timezone_id) references timezones (timezone_id)
);

/*
 * Данный sql запрос создает новую запись в таблице 'organizations_settings'.
 */
insert into organizations_settings (
    organization_id,
    country_id,
    location_id,
    organization_setting_address,
    organization_setting_postal_code,
    organization_setting_work_time,
    organization_setting_privacy,
    timezone_id
)
values (
    :organization_id,
    :country_id,
    :location_id,
    :organization_setting_address,
    :organization_setting_postal_code
    :organization_setting_work_time,
    :organization_setting_privacy,
    :timezone_id
)
returning
    organization_id;

/*
 * Данный sql запрос удаляет запись в таблице 'organizations_settings'.
 */
update
    organizations_settings
set
    entry_deleted_date_time = now()
where
    organization_id = $1;

/*
 * Данный sql запрос запрашивает запись из таблицы 'organizations_settings'.
 */
select
    organization_id,
    country_id,
    location_id,
    organization_setting_address,
    organization_setting_postal_code,
    organization_setting_work_time,
    organization_setting_privacy,
    timezone_id
from
    organizations_settings
where
    entry_deleted_date_time = null
and
    organization_id = $1;
/*
 * Данный запрос обновляет запись в таблице 'organizations_settings'
 */
update
    organizations_settings
set
    entry_updated_date_time = now(),
    %s
where
    organization_id = :organization_id
sreturning
    organization_id,
    country_id,
    location_id,
    organization_setting_address,
    organization_setting_postal_code,
    organization_setting_work_time,
    organization_setting_privacy,
    timezone_id;