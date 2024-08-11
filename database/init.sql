USE m1337Shirts;

DROP TABLE IF EXISTS product_image;
DROP TABLE IF EXISTS product;
DROP TABLE IF EXISTS user_profile;
DROP TABLE IF EXISTS promo;

CREATE TABLE product (
	id int NOT NULL AUTO_INCREMENT,
    title tinytext,
    price decimal(6,2),
    pdesc text,
    
    PRIMARY KEY (ID)
);

CREATE TABLE product_image (
	productId int NOT NULL,
    filepath varchar(80) NOT NULL,
    displayOrder int NOT NULL,
    PRIMARY KEY (filepath),
    FOREIGN KEY (productId) REFERENCES product(id)
);

CREATE TABLE user_profile (
	id int NOT NULL AUTO_INCREMENT,
    username tinytext NOT NULL,
    hashed_password varchar(32) NOT NULL,
    title tinytext,
    udesc text,
    profile_picture varchar(80),
    PRIMARY KEY (id)
);

CREATE TABLE promo (
	promo_code varchar(40) NOT NULL,
    promo_action tinytext NOT NULL,
    promo_title tinytext NOT NULL,
    PRIMARY KEY (promo_code)
);

INSERT INTO product (title, price, pdesc) VALUES ("Days Since Last Timezone issue: -1", 39.99, "When do timezones not cause problems");
INSERT INTO product (title, price, pdesc) VALUES ("It works on my machine", 29.99, "Sounds like a skill issue");
INSERT INTO product (title, price, pdesc) VALUES ("1 + 1 = 10", 34.99, "Binary, am I right. Used in so many things, and most people don't have slightiest idea how it works");
INSERT INTO product (title, price, pdesc) VALUES ("Stop bugging me", 29.99, "I have enough bugs in my program thank you very much");
INSERT INTO product (title, price, pdesc) VALUES ("Coffee Refill Drink", 29.99, "Ahh, the life source of any programmer. Can't get too far without a trusty cup of joe. I would many things for even a sip");
INSERT INTO product (title, price, pdesc) VALUES ("BigO Vibes", 39.99, "It's very satifying when you get that O(log n) time. I'd also be careful of anyone who's ok with O(n!)");
INSERT INTO product (title, price, pdesc) VALUES ("Trying to make things idiot-proof but they keep making better idiots", 24.99, "No matter how hard someone tries to make your program as user friendly as possible, you'll always get a person who just can't figure it out. If only we could get rid of those people right, haha...");
INSERT INTO product (title, price, pdesc) VALUES ("Roses are red, violets are blue, unexpected '{' on line 32", 34.99, "Error messages are a funny thing. Some languages have great error messages, but others couldn't be harder to understand (looking at you C++)");
INSERT INTO product (title, price, pdesc) VALUES ("Full Regex Cheatsheet", 29.99, "Boom, full regex cheetsheet. We already know that every time you write a regex is everytime you learn regex, my as well commit to the shirt now");

INSERT INTO product_image VALUES (1, "/media/products/daysSinceLastTimezone-1_0.jpg", 0);
INSERT INTO product_image VALUES (1, "/media/products/daysSinceLastTimezone-1_1.jpg", 1);
INSERT INTO product_image VALUES (2, "/media/products/itWorksOnMyMachine_0.jpg", 0);
INSERT INTO product_image VALUES (2, "/media/products/itWorksOnMyMachine_1.jpg", 1);
INSERT INTO product_image VALUES (3, "/media/products/1p1Equal10_0.jpg", 0);
INSERT INTO product_image VALUES (3, "/media/products/1p1Equal10_1.jpg", 1);
INSERT INTO product_image VALUES (4, "/media/products/stopBuggingMe_0.jpg", 0);
INSERT INTO product_image VALUES (4, "/media/products/stopBuggingMe_1.jpg", 1);
INSERT INTO product_image VALUES (5, "/media/products/coffeeRefillDrink_0.jpg", 0);
INSERT INTO product_image VALUES (5, "/media/products/coffeeRefillDrink_1.jpg", 1);
INSERT INTO product_image VALUES (6, "/media/products/bigOVibes_0.jpg", 0);
INSERT INTO product_image VALUES (6, "/media/products/bigOVibes_1.jpg", 1);
INSERT INTO product_image VALUES (7, "/media/products/idiotProofVsBetterIdiots_0.jpg", 0);
INSERT INTO product_image VALUES (7, "/media/products/idiotProofVsBetterIdiots_1.jpg", 1);
INSERT INTO product_image VALUES (8, "/media/products/redBlueLine32_0.jpg", 0);
INSERT INTO product_image VALUES (8, "/media/products/redBlueLine32_1.jpg", 1);
INSERT INTO product_image VALUES (9, "/media/products/fullRegexCheatSheet_0.jpg", 0);
INSERT INTO product_image VALUES (9, "/media/products/fullRegexCheatSheet_1.jpg", 1);

# important users
INSERT INTO user_profile (username, hashed_password, udesc, title, profile_picture)
VALUES ("tednugent", "ab148bc57795c415b8dc46bc4ba01744", "I love music as much as the next guy, but programming humor is what I really love. I do lots of it, it's so fun.", "partner", "/media/profiles/TedNugent.jpg");
INSERT INTO user_profile (username, hashed_password, udesc, title, profile_picture)
VALUES ("tomselleck", "1c2bb845ed42ae28469d362bd8889ffe", "I've done a number of movies in my times, but most people don't notice how often I program, so I've created this page to spread my joy of programming.", "partner", "/media/profiles/TomSelleck.jpg");
INSERT INTO user_profile (username, hashed_password, udesc, title, profile_picture)
VALUES ("tracymarrow", "2d08a583daa61da0452fc1a1de38f1fe", "Programming my whole life and it's always a good time, you know. I got nothing without my keyboard.", "partner", "/media/profiles/TracyMarrow.jpg");

# filler users
INSERT INTO user_profile (username, hashed_password, udesc, title)
VALUES ('alice_wonder', '7d793037a0770186574b0282f2f435e7', 'Hey there, I am Alice. I have a knack for solving complex problems and love to share my knowledge.', 'employee');
INSERT INTO user_profile (username, hashed_password, udesc, title)
VALUES ('bob_miller', '1d2292f5d6b12e0c3e57d52fb5ef7922', 'Hi, I am Bob. I am known for my constructive contributions and detailed code reviews.', 'contributor');
INSERT INTO user_profile (username, hashed_password, udesc, title)
VALUES ('charlie_jones', '8f14e45fceaa167a5a36dedd4bea2543', 'Hello, I am Charlie. I enjoy working on front-end development and advocate for clean code.', 'employee');
INSERT INTO user_profile (username, hashed_password, udesc, title)
VALUES ('david_black', 'e4da3b7fbbcef345d7772b0674a318d5', 'Hi, I am David. I am a backend specialist with extensive experience in database management.', 'contributor');
INSERT INTO user_profile (username, hashed_password, udesc, title)
VALUES ('eve_white', '1679091c5a880fef6fb5e6087eb1b2dc', 'Hey, I am Eve. I am an all-rounder with a deep understanding of full-stack development.', 'employee');
INSERT INTO user_profile (username, hashed_password, udesc, title)
VALUES ('jane_smith', '5d41402abc4b2a76b9719d911017c592', 'Hello, I am Jane. I am a passionate contributor to various coding communities.', 'contributor');
INSERT INTO user_profile (username, hashed_password, udesc, title)
VALUES ('frank_green', '45c48cce2e2d3fbdea1afc51c7c6ad26', 'Hi, I am Frank. I am a seasoned developer who enjoys mentoring junior team members.', 'contributor');
INSERT INTO user_profile (username, hashed_password, udesc, title)
VALUES ('grace_lee', '03c7c0ace395d82182db07ae2c30f034', 'Hello, I am Grace. I have a long history of contributions to the tech world.', 'employee');
INSERT INTO user_profile (username, hashed_password, udesc, title)
VALUES ('john_roe', 'e99a18c428cb38d5f260853678922e03', 'Hi, I am John. I am an experienced software developer who enjoys working on open-source projects.', 'employee');
INSERT INTO user_profile (username, hashed_password, udesc, title)
VALUES ('henry_clark', 'c4ca4238a0b943820dcc509a6f75849b', 'Hi, I am Henry. I am skilled at finding and fixing bugs quickly and efficiently.', 'contributor');

# Inside man
INSERT INTO user_profile (username, hashed_password, udesc, title)
VALUES ("CM", "286ba48ca4f45dab999e3609fcf562dc", "You will recieve instructions from me here<br>https://youtube.com", "hidden");

INSERT INTO promo VALUES ("minus10", "DISCOUNT 10", "10% Discount");
INSERT INTO promo VALUES ("minus50", "DISCOUNT 50", "50% Discount");
INSERT INTO promo VALUES ("truth{all_your_db_is_belong_to_us}", "REVEAL TRUTH", "Revealing the truth");

CREATE DATABASE m1337Guns;
USE m1337Guns;

DROP TABLE IF EXISTS product;

CREATE TABLE product (
	id int NOT NULL AUTO_INCREMENT,
    title tinytext,
    price decimal(6,2),
    pdesc text,
    imagepath varchar(80),
    PRIMARY KEY (ID)
);

INSERT INTO product (title, price, pdesc) VALUES ("Ak-47", 1199.99, "<strong>Rate of fire</strong><br>
Cyclic rate: 600 rounds/min Practical rate: Semi-automatic: 40 rounds/min Bursts/ Fully automatic: 100 rounds/min<br>
<strong>Muzzle velocity</strong><br>
715 m/s (2,350 ft/s)<br>
<strong>Effective firing range</strong><br>
350 m (380 yd)<br>
<strong>Feed system</strong><br>
20-round, 30-round, 50-round detachable box magazine, 40-round, 75-round drum magazines also available<br>",
"/static/media/products/AK47.jpg");
INSERT INTO product (title, price, pdesc) VALUES ("Winchester Model 1911 Self Loading shotgun", 899.99, "<strong>Caliber</strong><br>
12 gauge, 16 gauge, 20 gauge, 28 gauge<br>
<strong>Action</strong><br>
Long recoil<br>
<strong>Feed system</strong><br>
5-round tubular magazine<br>
<strong>Caliber</strong><br>
Bead<br>", "/static/media/products/SelfLoadingShotgun.jpg");
INSERT INTO product (title, price, pdesc) VALUES ("M240 Machine Gun", 8499.99, "<strong>Muzzle velocity</strong><br>
2,800 ft/s (853 m/s)<br>
<strong>Effective firing range</strong><br>
800–1,800 m (875–1,969 yd) depending on mount<br>
<strong>Maximum firing range</strong><br>
3,725 m (4,074 yd)<br>
<strong>Feed system</strong><br>
Belt-fed using M13 disintegrating links, 50-round ammo pouch, or non-disintegrating DM1 belt<br>",
"/static/media/products/M240MG.png");
INSERT INTO product (title, price, pdesc) VALUES ("HK416", 3499.99, "<strong>Barrel length</strong><br>
11–20 in (280–510 mm) HK416C: 9 in (230 mm)<br>
<strong>Width</strong><br>
78 mm (3.1 in)<br>
<strong>Height</strong><br>
236–240 mm (9.3–9.4 in)
<strong>Cartridge</strong><br>
5.56×45mm NATO<br>", "/static/media/products/HK416N.png");
INSERT INTO product (title, price, pdesc) VALUES ("M4 Carbine", 599.99, "<strong>Weight</strong><br>
6.36 lbs<br>
<strong>Length</strong><br>
33 in (stock extended)<br>
<strong>Barrel length</strong><br>
14.5 in
<strong>Caliber</strong><br>
5.56x45 mm<br>", "/static/media/products/M4Carbine.webp");
INSERT INTO product (title, price, pdesc) VALUES ("Uzi", 5499.99, "<strong>Weight</strong><br>
7.72 lb<br>
<strong>Length</strong><br>
445 mm (17.5 in) stockless 470 mm (18.5 in) folding stock collapsed 640 mm (25 in) folding stock extended<br>
<strong>Barrel length</strong><br>
260 mm (10.2 in)
<strong>Cartridge</strong><br>
.22 LR .41 AE .45 ACP 9×19mm Parabellum 9×21mm IMI<br>", "/static/media/products/uzi.jpg");
INSERT INTO product (title, price, pdesc) VALUES ("RPG-29", 5499.99, "<strong>Mass</strong><br>
12.1 kg (27 lb) unloaded (with optical sight) 18.8 kg (41 lb) loaded (ready to fire)<br>
<strong>Length</strong><br>
1 m (3 ft 3 in) (dismantled for transport) 1.85 m (6 ft 1 in) (ready to fire)<br>
<strong>Cartridge</strong><br>
PG-29V tandem rocket TBG-29V thermobaric rounds
<strong>Caliber</strong><br>
105 mm (4.1 in) barrel 65 and 105 mm (2.6 and 4.1 in) warheads<br>", "/static/media/products/RPG29.jpg");
INSERT INTO product (title, price, pdesc) VALUES ("Desert Eagle", 5499.99, "<strong>Mass</strong><br>
4.4 lb<br>
<strong>Length</strong><br>
14.75 in (10 in barrel)<br>
<strong>Action</strong><br>
Gas-operated, closed rotating bolt
<strong>Muzzle velocity</strong><br>
1542 ft/s (470 m/s) (.50 AE)<br>", "/static/media/products/desertEagle.png");
INSERT INTO product (title, price, pdesc, filepath) VALUES ("McMillan TAC-50C", 11499.99, "<strong>Caliber</strong><br>
.50 BMG<br>
<strong>Barrel</strong><br>
Match Grade, Stainless Steel<br>
<strong>Length</strong><br>
56.5 in
<strong>Weight</strong><br>
29 lbs<br>", "/media/products/TAC-50C.jpg");

USE m1337Shirts;