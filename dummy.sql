drop database project_ecommerce_k2;
create database project_ecommerce_k2;
use project_ecommerce_k2;

-- user Dummy
INSERT INTO users (name, email, password, status) VALUES ("admin", "admin@gmail.com", "admin", 1);
INSERT INTO users (name, email, password, status) VALUES ("testuser1", "testuser1@gmail.com", "testuser1", -1);
INSERT INTO users (name, email, password, status) VALUES ("testuser2", "testuser2@gmail.com", "testuser2", -1);
INSERT INTO users (name, email, password, status) VALUES ("testuser3", "testuser3@gmail.com", "testuser3", -1);
INSERT INTO users (name, email, password, status) VALUES ("testuser4", "testuser4@gmail.com", "testuser4", -1);
--

INSERT INTO addresses (street, city, region, postal_code, user_id) VALUES ("teststreetaddres1", "city1", "region1", "postal_code1",1);
INSERT INTO addresses (street, city, region, postal_code, user_id) VALUES ("teststreetaddres2", "city2", "region2", "postal_code2",2);
INSERT INTO addresses (street, city, region, postal_code, user_id) VALUES ("teststreetaddres3", "city3", "region3", "postal_code3",3);


INSERT INTO categories (name) VALUES ("category1");
INSERT INTO categories (name) VALUES ("category2");
INSERT INTO categories (name) VALUES ("category3");
INSERT INTO categories (name) VALUES ("category4");



INSERT INTO products (name, price, qty, description ,category_id) VALUES ("testerproductname1", 20000, 10, "testerproductdescription1",1);
INSERT INTO products (name, price, qty, description ,category_id) VALUES ("testerproductname2", 450000, 20, "testerproductdescription2",1);
INSERT INTO products (name, price, qty, description ,category_id) VALUES ("testerproductname3", 220000, 30, "testerproductdescription3",3);
INSERT INTO products (name, price, qty, description ,category_id) VALUES ("testerproductname4", 40000, 40, "testerproductdescription4",4);
INSERT INTO products (name, price, qty, description ,category_id) VALUES ("testerproductname5", 60000, 30, "testerproductdescription5",3);




	
	-- mockUser := entities.User{Name: "admin", Email: "admin@gmail.com", Password: "testuser1", Status: 1}
	-- repoUser.Register(mockUser)
	-- mockUser1 := entities.User{Name: "testuser1", Email: "testuser1@gmail.com", Password: "testuser1", Status: -1}
	-- repoUser.Register(mockUser1)
	-- mockUser2 := entities.User{Name: "testuser2", Email: "testuser2@gmail.com", Password: "testuser2", Status: -1}
	-- repoUser.Register(mockUser2)
	-- mockUser3 := entities.User{Name: "testuser3", Email: "testuser3@gmail.com", Password: "testuser3", Status: -1}
	-- repoUser.Register(mockUser3)
	-- //

	//Address Dummy
	mockAddress1 := entities.Address{Street: "teststreetaddres1", City: "city1", Region: "region1", Postal_code: "postal_code1", User_ID: 1}
	repoAddress.Insert(mockAddress1)
	mockAddress2 := entities.Address{Street: "teststreetaddres2", City: "city2", Region: "region2", Postal_code: "postal_code2", User_ID: 2}
	repoAddress.Insert(mockAddress2)
	mockAddress3 := entities.Address{Street: "teststreetaddres3", City: "city3", Region: "region3", Postal_code: "postal_code3", User_ID: 3}
	repoAddress.Insert(mockAddress3)
	//

	//Category Dummy
	// mockCategory1 := entities.Category{Name: "category1"}
	// repoCategory.Create(mockCategory1)
