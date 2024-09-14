Begin;

INSERT INTO `clients` (`name`, `api_key`, `is_active`, `created_at`, `updated_at`)
VALUES
    ('core', 'OXKDaaavd7UAMmNf2LcsXz13mcdddFS7UYlE87ld39TcmsQCabccc7jhL4YCPzk', 1, now(), now()),
ON DUPLICATE KEY UPDATE `updated_at` = now();

commit;