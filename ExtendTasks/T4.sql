-- MariaDB dump 10.19  Distrib 10.6.5-MariaDB, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: S_T_U201911665
-- ------------------------------------------------------
-- Server version	10.6.5-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `Course`
--

DROP TABLE IF EXISTS `Course`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Course` (
  `Cno` char(4) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '课程号',
  `Cname` char(40) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '课程名',
  `Cpno` char(4) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '先行课',
  `Ccredit` smallint(6) NOT NULL COMMENT '学分',
  PRIMARY KEY (`Cno`),
  KEY `Cpno_fk` (`Cpno`),
  CONSTRAINT `Cpno_fk` FOREIGN KEY (`Cpno`) REFERENCES `Course` (`Cno`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Course`
--

LOCK TABLES `Course` WRITE;
/*!40000 ALTER TABLE `Course` DISABLE KEYS */;
INSERT INTO `Course` VALUES ('1','数据库','5',4),('2','数学',NULL,2),('3','信息系统','1',4),('4','操作系统','6',3),('5','数据结构','7',4),('6','数据处理',NULL,2),('7','java','6',4);
/*!40000 ALTER TABLE `Course` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Temporary table structure for view `CourseGrade`
--

DROP TABLE IF EXISTS `CourseGrade`;
/*!50001 DROP VIEW IF EXISTS `CourseGrade`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE TABLE `CourseGrade` (
  `Cno` tinyint NOT NULL,
  `Cname` tinyint NOT NULL,
  `Cpno` tinyint NOT NULL,
  `Ccredit` tinyint NOT NULL,
  `Grade` tinyint NOT NULL,
  `Sno` tinyint NOT NULL
) ENGINE=MyISAM */;
SET character_set_client = @saved_cs_client;

--
-- Table structure for table `SC`
--

DROP TABLE IF EXISTS `SC`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `SC` (
  `Sno` char(9) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'StudentNumber',
  `Cno` char(4) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'ClassNumber',
  `Grade` smallint(6) DEFAULT NULL COMMENT 'Performance in course',
  PRIMARY KEY (`Sno`,`Cno`),
  KEY `Cno_index` (`Cno`),
  KEY `Sno_index` (`Sno`),
  CONSTRAINT `Cno_fk` FOREIGN KEY (`Cno`) REFERENCES `Course` (`Cno`),
  CONSTRAINT `Sno_fk` FOREIGN KEY (`Sno`) REFERENCES `Student` (`Sno`),
  CONSTRAINT `c_GradeRange` CHECK (`Grade` between 0 and 100)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='选修关系';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `SC`
--

LOCK TABLES `SC` WRITE;
/*!40000 ALTER TABLE `SC` DISABLE KEYS */;
INSERT INTO `SC` VALUES ('200215121','1',92),('200215121','2',85),('200215121','3',88),('200215122','2',90),('200215122','3',80);
/*!40000 ALTER TABLE `SC` ENABLE KEYS */;
UNLOCK TABLES;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_unicode_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'IGNORE_SPACE,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`localhost`*/ /*!50003 TRIGGER T_scholar
    before update
    on SC
    for each row
    IF NEW.Grade > 95 THEN
        IF EXISTS(SELECT * from Student where Student.Sno = OLD.Sno and Student.Scholarship = '否') THEN
            UPDATE Student SET Scholarship = '是' where Student.Sno = OLD.Sno;
        end if;
    ELSE
        IF EXISTS(SELECT * from SC where SC.Sno = OLD.Sno and SC.Grade <= 95) THEN
            IF OLD.Grade > 95 THEN
                UPDATE Student SET Scholarship = '否' where Student.Sno = OLD.Sno;
            end if;
        end if;
    end if */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;

--
-- Table structure for table `Student`
--

DROP TABLE IF EXISTS `Student`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Student` (
  `Sno` char(9) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '学号',
  `Sname` char(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '姓名',
  `Ssex` char(2) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '性别',
  `Sage` smallint(6) NOT NULL COMMENT '年龄',
  `Sdept` char(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '所在系',
  `Scholarship` char(2) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '否' COMMENT '奖学金',
  PRIMARY KEY (`Sno`),
  UNIQUE KEY `Student_Sname_uindex` (`Sname`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Student`
--

LOCK TABLES `Student` WRITE;
/*!40000 ALTER TABLE `Student` DISABLE KEYS */;
INSERT INTO `Student` VALUES ('200215121','李勇','男',20,'CS','否'),('200215122','刘晨','女',19,'CS','否'),('200215123','王敏','女',18,'MA','否'),('200215125','张立','男',19,'CS','是');
/*!40000 ALTER TABLE `Student` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Temporary table structure for view `StudentGrade`
--

DROP TABLE IF EXISTS `StudentGrade`;
/*!50001 DROP VIEW IF EXISTS `StudentGrade`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE TABLE `StudentGrade` (
  `Sno` tinyint NOT NULL,
  `Sname` tinyint NOT NULL,
  `Ssex` tinyint NOT NULL,
  `Sage` tinyint NOT NULL,
  `Scholarship` tinyint NOT NULL,
  `Grade` tinyint NOT NULL,
  `Cno` tinyint NOT NULL,
  `Sdept` tinyint NOT NULL
) ENGINE=MyISAM */;
SET character_set_client = @saved_cs_client;

--
-- Dumping events for database 'S_T_U201911665'
--

--
-- Dumping routines for database 'S_T_U201911665'
--
/*!50003 DROP PROCEDURE IF EXISTS `DeleteCourseStrict` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_unicode_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'IGNORE_SPACE,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `DeleteCourseStrict`()
BEGIN
    DELETE
    FROM Course
    where Cno in
          (SELECT Cno
           from Course as C1
           where Cno not in (SELECT SC.Cno from SC where SC.Cno = C1.Cno)
             and Cno not in (SELECT C2.Cpno from Course as C2 where C2.Cpno = C1.Cno));
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `ShowDeptGrade` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_unicode_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'IGNORE_SPACE,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `ShowDeptGrade`(IN deptName varchar(20), IN CourseNumber varchar(4))
BEGIN
    SELECT AVG(Grade) as avg,MAX(Grade) as max,MIN(Grade) as min,SUM(IF(Grade >= 90,1,0))/COUNT(*) as rate, SUM(IF(Grade < 60,1,0)) as fail from SC where Sno in (SELECT Sno FROM Student WHERE Sdept = deptName) and Cno = CourseNumber and Grade is not null;
end ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;

--
-- Final view structure for view `CourseGrade`
--

/*!50001 DROP TABLE IF EXISTS `CourseGrade`*/;
/*!50001 DROP VIEW IF EXISTS `CourseGrade`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_unicode_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `CourseGrade` AS select `Course`.`Cno` AS `Cno`,`Course`.`Cname` AS `Cname`,`Course`.`Cpno` AS `Cpno`,`Course`.`Ccredit` AS `Ccredit`,`SC`.`Grade` AS `Grade`,`SC`.`Sno` AS `Sno` from (`SC` join `Course`) where `SC`.`Cno` = `Course`.`Cno` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `StudentGrade`
--

/*!50001 DROP TABLE IF EXISTS `StudentGrade`*/;
/*!50001 DROP VIEW IF EXISTS `StudentGrade`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_unicode_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `StudentGrade` AS select `Student`.`Sno` AS `Sno`,`Student`.`Sname` AS `Sname`,`Student`.`Ssex` AS `Ssex`,`Student`.`Sage` AS `Sage`,`Student`.`Scholarship` AS `Scholarship`,`SC`.`Grade` AS `Grade`,`SC`.`Cno` AS `Cno`,`Student`.`Sdept` AS `Sdept` from (`Student` join `SC`) where `SC`.`Sno` = `Student`.`Sno` order by `SC`.`Grade` desc */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-12-25 16:37:28
