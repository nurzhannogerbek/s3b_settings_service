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
    organization_id uuid not null,
    foreign key (organization_id) references organizations (organization_id),
    country_id uuid not null,
    foreign key (country_id) references countries (country_id),
    location_id uuid not null,
    foreign key (location_id) references locations (location_id),
    organization_setting_address text null,
    organization_setting_postal_code varchar (50) null,
    organization_setting_work_time text not null,
    organization_setting_privacy text not null,
    timezone_id uuid not null,
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
    :organization_setting_privacy
    :timezone_id
)
returning
    organization_id;