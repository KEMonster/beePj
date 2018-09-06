CREATE TABLE `tb_user_token`(
    `iId` int(10) NOT NULL AUTO_INCREMENT,
    `sToken` varchar(50) NULL DEFAULT '',
    `sKey` varchar(128) NULL DEFAULT '',
    `sUserCode` varchar(20) NULL DEFAULT '',
    PRIMARY KEY(`iId`),
    FULLTEXT KEY `idx_user_token` (`sToken`,`sUserCode`) 
)ENGINE=MyISAM DEFAULT charset=UTF8;


CREATE TABLE `tb_order`(
    `sOrderCode` int(10) NOT NULL AUTO_INCREMENT,
    `sMethod` varchar(20) NULL DEFAULT '',
    `sZipCode` varchar(10) NULL DEFAULT '',
    `sSendName` varchar(50) NULL DEFAULT '',
    `sSendAddress` varchar(128) NULL DEFAULT '',
    `sRecipientName` varchar(50) NULL DEFAULT '',
    `sRecipientAddress` varchar(128) NULL DEFAULT '',
    `sUserCode` varchar(20) NOT NULL DEFAULT '',
    PRIMARY KEY(`sOrderCode`),
    KEY `idx_usercode` (`sUserCode`),
    KEY `idx_method` (`sMethod`) 
)ENGINE=Innodb DEFAULT charset=UTF8;
