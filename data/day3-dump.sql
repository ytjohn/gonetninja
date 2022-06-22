PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE activity
(
    id          uuid      not null
        constraint activity_pk
            primary key,
    created_at  timestamp not null,
    updated_at  timestamp not null,
    net         uuid      not null
        constraint activity_nets_id_fk
            references nets,
    action      text      not null,
    time_at     timestamp not null,
    name        text      not null,
    description text
);
INSERT INTO activity VALUES('82a6d9f6-fce2-4e4d-9267-abd263c1201e',1655894025000,1655894028000,'786829ca-f1c3-11ec-a9d5-ce5390833a0f','netcontrol',1655893740000,'n0call','assume net control');
INSERT INTO activity VALUES('56f06fb9-44df-4ee1-9d1f-819a02e20d40',1655894025000,1655894028000,'786829ca-f1c3-11ec-a9d5-ce5390833a0f','checkin',1655893747000,'n1call','early checkin');
INSERT INTO activity VALUES('55273830-2eae-409b-ab14-2f07d37873b9',1655894025000,1655894028000,'786829ca-f1c3-11ec-a9d5-ce5390833a0f','checkin',1655893747000,'n2call','early checkin');
INSERT INTO activity VALUES('55273830-2eae-409b-ab14-2f07d37873b7',1655894025000,1655894028000,'786829ca-f1c3-11ec-a9d5-ce5390833a0f','checkin',1655894047000,'n3call','regular checkin');
INSERT INTO activity VALUES('a3c329e6-1457-4c6b-a358-5258528fadc8',1655894025000,1655894028000,'786829ca-f1c3-11ec-a9d5-ce5390833a0f','checkin',1655894047000,'n4call','regular checkin');
INSERT INTO activity VALUES('bd311989-0e73-44a6-8224-96ca61de1df4',1655894025000,1655894028000,'786829ca-f1c3-11ec-a9d5-ce5390833a0f','close',1655894227000,'n5call','close net');
INSERT INTO activity VALUES('99cf7ae7-3287-439a-bf5d-c8767fd7d4c9',1655894025000,1655894028000,'786829ca-f1c3-11ec-a9d5-ce5390833a0f','open',1655893869000,'n0call','open net');
INSERT INTO activity VALUES('04b881db-68a1-49d1-b6c6-446a775e5e2c',1655894025000,1655894028000,'5751c33c-f230-11ec-a9d5-ce5390833a0f','open',1655893987000,'n0call','open net');
INSERT INTO activity VALUES('d43423dc-9d51-4046-8de5-63490dccd33c',1655894025000,1655894028000,'786829ca-f1c3-11ec-a9d5-ce5390833a0f','checkout',1655893760000,'n2call','one and done');
INSERT INTO activity VALUES('5d98eb39-453c-4301-a656-43e8b32ead60',1655894025000,1655894028000,'786829ca-f1c3-11ec-a9d5-ce5390833a0f','comment',1655894057000,'n4call','i need to know who is getting chicken dinner');
INSERT INTO activity VALUES('10648876-7e95-4f3c-8eec-382635addc32',1655894025000,1655894028000,'786829ca-f1c3-11ec-a9d5-ce5390833a0f','netcontrol',1655894049000,'n5call','assume net control');
INSERT INTO activity VALUES('362d6f2a-6257-4ba9-be5f-049f6acd750c',1655894025000,1655894028000,'786829ca-f1c3-11ec-a9d5-ce5390833a0f','checkin',1655894407000,'n6call','tried to check in earlier, but could not make repeater');
CREATE TABLE IF NOT EXISTS "nets"
(
    id            uuid      not null
        constraint nets_pk
            primary key,
    created_at    timestamp not null,
    updated_at    timestamp not null,
    name          text      not null,
    planned_start timestamp not null,
    planned_end   timestamp not null
);
INSERT INTO nets VALUES('c31911fc-5d3c-4b18-b4b1-1e081aa6effd',1655724491000,1655724491000,'test one',1655724491000,1655724491000);
INSERT INTO nets VALUES('ccad3aad-c9ea-4891-a604-8d02e0968ce8',1655724491000,1655724491000,'test two',1655724491000,1655724491000);
INSERT INTO nets VALUES('66b685a9-ea20-4a14-b766-4d23f362be4b',1655724491000,1655724491000,'test three',1655724491000,1655724491000);
INSERT INTO nets VALUES('2bc9c10a-4056-45bc-bbb6-6a482bbb30a9',1655724491000,1655724491000,'test the fourth',1655724491000,1655724491000);
INSERT INTO nets VALUES('2a590052-639c-4db9-afdc-343de5accba7',1655724491000,1655724491000,'Daily Net 5',1655724491000,1655724491000);
INSERT INTO nets VALUES('842fd70a-f1bc-11ec-9f5f-ce5390833a0f',1655724491000,1655724491000,'Tuesday Night Net',1655724491000,1655724491000);
INSERT INTO nets VALUES('33780dd0-f1be-11ec-8bd4-ce5390833a0f',1655724491000,1655724491000,'dfasdf',1655724491000,1655724491000);
INSERT INTO nets VALUES('584ea790-f1be-11ec-8bd4-ce5390833a0f',1655724491000,1655724491000,'asfsd',1655724491000,1655724491000);
INSERT INTO nets VALUES('786829ca-f1c3-11ec-a9d5-ce5390833a0f',1655724491000,1655724491000,'Tuesday night june 21st net',1655724491000,1655724491000);
INSERT INTO nets VALUES('5751c33c-f230-11ec-a9d5-ce5390833a0f',1655724491000,1655724491000,'Rocking new net for screenshot',1655724491000,1655724491000);
INSERT INTO nets VALUES('adda0132-f240-11ec-b7d0-ce5390833a0f','2022-06-22 11:33:35.02201-04:00','2022-06-22 11:33:35.02201-04:00','foo 300','2022-06-22 12:00:00+00:00','2022-06-22 12:30:00+00:00');
INSERT INTO nets VALUES('7bce5430-f273-11ec-96c4-ce5390833a0f','2022-06-22 17:37:15.392967-04:00','2022-06-22 17:37:15.392967-04:00','High Noon HW Net','2022-06-22 12:00:00+00:00','2022-06-22 12:30:00+00:00');
CREATE UNIQUE INDEX activity_id_uindex
    on activity (id);
CREATE INDEX activity_name_index
    on activity (name);
COMMIT;
