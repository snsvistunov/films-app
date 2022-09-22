CREATE TABLE "user"
(
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY, 
    login varchar(255) not null, 
    password varchar(255) not null, 
    age smallint not null
);

CREATE TABLE roles
(
    id serial not null unique PRIMARY KEY, 
    name varchar(255) not null, 
    description text not null 
);

CREATE TABLE user_roles
(
    user_id uuid references "user"(id) on delete cascade not null, 
    roles_id smallint references roles(id) on delete cascade not null  
);

CREATE TABLE director
(
   id uuid DEFAULT gen_random_uuid() PRIMARY KEY,  
   name varchar(255) not null, 
   date_of_birth date not null,
   photo varchar(255) not null 
);

CREATE TABLE film
(
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY, 
    name varchar(255) not null, 
    genre varchar(255) not null, 
    director_id uuid references director(id) on delete cascade not null, 
    rate numeric (3,1) not null,
    year date not null,
    minutes smallint not null
);

CREATE TABLE wishlist
(
    user_id uuid references "user"(id) on delete cascade not null, 
    film_id uuid references film(id) on delete cascade not null  
);

CREATE TABLE favourites
(
    user_id uuid references "user"(id) on delete cascade not null, 
    film_id uuid references film(id) on delete cascade not null  
);