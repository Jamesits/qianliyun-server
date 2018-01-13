BEGIN TRANSACTION;

CREATE TABLE userInfo (ID INTEGER PRIMARY KEY, Username TEXT UNIQUE NOT NULL, Password TEXT, Salt TEXT, Alias TEXT, ResellerAlias TEXT, AuthMax INTEGER, AuthLeft INTEGER, DeauthLeft INTEGER, Reseller INTEGER);
INSERT INTO userInfo VALUES(1,'root',NULL,NULL,NULL,NULL,2147483647,2147483647,2147483647,NULL);
CREATE TABLE liveSession (ID INTEGER PRIMARY KEY, UserID INTEGER NOT NULL, URL TEXT, Title TEXT, Host TEXT, Comment TEXT, Begin REAL, End REAL, Tags TEXT);
CREATE TABLE customerInfo (ID INTEGER PRIMARY KEY, UserID INTEGER NOT NULL, CustomerName TEXT, Mobile TEXT, Status TEXT, Tags TEXT);
CREATE TABLE liveActivity (ID INTEGER PRIMARY KEY, UserID INTEGER NOT NULL, LiveID INTEGER, Time REAL, CustomerID INTEGER, Activity TEXT);

INSERT INTO liveSession (ID, UserID, URL, Title, Tags) VALUES (1, 1, 'bb4b7f4f-828f-45f3-93eb-4ddb47979376', '测试1', '女装, 箱包');
INSERT INTO liveSession (ID, UserID, URL, Title, Tags) VALUES (2, 1, '6d7c46fc-4661-453b-92dd-cad8ca3a2965', '测试2', '数码');
INSERT INTO customerInfo (ID, UserID, CustomerName, Mobile, Status, Tags) VALUES (1, 1, '小张', '+86 138-0013-8000', '已打标', '女装, 箱包');
INSERT INTO customerInfo (ID, UserID, CustomerName, Mobile, Status, Tags) VALUES (2, 1, '小李', '+86 139-1234-5678', '已打标', '数码, 女装');
INSERT INTO customerInfo (ID, UserID, CustomerName, Mobile, Status, Tags) VALUES (3, 1, '小王', '+86 139-1234-5678', '已拒绝', '数码');
INSERT INTO liveActivity (ID, UserID, LiveID, CustomerID, Activity) VALUES (1, 1, 1, 1, '进入了直播间');
INSERT INTO liveActivity (ID, UserID, LiveID, CustomerID, Activity) VALUES (2, 1, 1, 2, '进入了直播间');
INSERT INTO liveActivity (ID, UserID, LiveID, CustomerID, Activity) VALUES (3, 1, 2, 2, '进入了直播间');
INSERT INTO liveActivity (ID, UserID, LiveID, CustomerID, Activity) VALUES (4, 1, 2, 3, '进入了直播间');

COMMIT;
