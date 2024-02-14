create table currency(
    Id serial primary key,
    Date date not null,
    Time_stamp 
    Base character varying(3),
    Rate character varying(3),
    Value numeric(15,13)
)