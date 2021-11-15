
3create table pathways(
	pathway_id serial not null primary key,
	name varchar(128) not null
);

create table courses(
	course_id serial not null primary key,
	name varchar(128) not null,
	is_active boolean default false,
	created_at  timestamp with time zone default current_timestamp,
	/*update*/
	pathway_id int not null references pathway(pathway_id)
);


create table takeaways(
	takeaway_id serial not null primary key,
	content varchar(256) not null,
	course_id int not null references courses(course_id)
);

create table chapters(
	chapter_id serial not null primary key,
	name varchar(128) not null,
	course_id int not null references courses(course_id)
);

create table lessons(
	lesson_id serial not null primary key,
	name varchar(128) not null,
	chapter_id int not null references chapters(chapter_id)
);
