REPLACE INTO urms.config (`key`, value) VALUES ('vmware.archivationWindowStart', '23:00:00');
REPLACE INTO urms.config (`key`, value) VALUES ('vmware.archivationWindowEnd', '06:00:00');

DROP TABLE IF EXISTS `vmware_request`;
DROP TABLE IF EXISTS `vmware_tasks`;
CREATE TABLE `vmware_tasks` ( `ID` int  NOT NULL auto_increment,
    `action` char(11) NOT NULL, `vmkey` char(10) NOT NULL DEFAULT '',
    `parameters` varchar(255) NOT NULL DEFAULT '', `status` char(11) NOT NULL DEFAULT '',
    PRIMARY KEY (`ID`)) ENGINE = InnoDB DEFAULT CHARSET = utf8;