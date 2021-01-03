CREATE DATABASE  IF NOT EXISTS `hummingbird` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `hummingbird`;
-- MySQL dump 10.13  Distrib 8.0.22, for Win64 (x86_64)
--
-- Host: localhost    Database: hummingbird
-- ------------------------------------------------------
-- Server version	8.0.22

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `badge`
--

DROP TABLE IF EXISTS `badge`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `badge` (
  `id` int NOT NULL AUTO_INCREMENT,
  `keyword` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `keyword_description` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_badge` (`keyword`,`keyword_description`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `badge`
--

LOCK TABLES `badge` WRITE;
/*!40000 ALTER TABLE `badge` DISABLE KEYS */;
/*!40000 ALTER TABLE `badge` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `category`
--

DROP TABLE IF EXISTS `category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `category` (
  `id` int NOT NULL AUTO_INCREMENT,
  `creator_id` int NOT NULL,
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `title` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `description` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `avatar` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `cover` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `color` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `order` int DEFAULT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_category_creator` (`creator_id`),
  CONSTRAINT `fk_category_creator` FOREIGN KEY (`creator_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category`
--

LOCK TABLES `category` WRITE;
/*!40000 ALTER TABLE `category` DISABLE KEYS */;
/*!40000 ALTER TABLE `category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `category_community`
--

DROP TABLE IF EXISTS `category_community`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `category_community` (
  `id` int NOT NULL AUTO_INCREMENT,
  `category_id` int NOT NULL,
  `community_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_category_community_category` (`category_id`),
  KEY `fk_category_community_community` (`community_id`),
  CONSTRAINT `fk_category_community_category` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`),
  CONSTRAINT `fk_category_community_community` FOREIGN KEY (`community_id`) REFERENCES `community` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category_community`
--

LOCK TABLES `category_community` WRITE;
/*!40000 ALTER TABLE `category_community` DISABLE KEYS */;
/*!40000 ALTER TABLE `category_community` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `circle`
--

DROP TABLE IF EXISTS `circle`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `circle` (
  `id` int NOT NULL AUTO_INCREMENT,
  `creator_id` int NOT NULL,
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `color` varchar(7) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `users_count` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_circle_creator` (`creator_id`),
  CONSTRAINT `fk_circle_creator` FOREIGN KEY (`creator_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `circle`
--

LOCK TABLES `circle` WRITE;
/*!40000 ALTER TABLE `circle` DISABLE KEYS */;
INSERT INTO `circle` VALUES (1,210,'Technology','#023ca7','2020-07-19 15:58:27',100);
/*!40000 ALTER TABLE `circle` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `circle_post`
--

DROP TABLE IF EXISTS `circle_post`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `circle_post` (
  `id` int NOT NULL AUTO_INCREMENT,
  `circle_id` int NOT NULL,
  `post_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_circle_post_circle` (`circle_id`),
  KEY `fk_circle_post_post` (`post_id`),
  CONSTRAINT `fk_circle_post_circle` FOREIGN KEY (`circle_id`) REFERENCES `circle` (`id`),
  CONSTRAINT `fk_circle_post_post` FOREIGN KEY (`post_id`) REFERENCES `post` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `circle_post`
--

LOCK TABLES `circle_post` WRITE;
/*!40000 ALTER TABLE `circle_post` DISABLE KEYS */;
INSERT INTO `circle_post` VALUES (1,1,201),(2,1,202),(3,1,203);
/*!40000 ALTER TABLE `circle_post` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `community`
--

DROP TABLE IF EXISTS `community`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `community` (
  `id` int NOT NULL AUTO_INCREMENT,
  `creator_id` int NOT NULL,
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `title` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `description` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `rules` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `avatar` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `color` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `type` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `user_adjective` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `users_adjective` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `invites_enabled` bit(1) NOT NULL DEFAULT (1),
  `is_deleted` bit(1) NOT NULL DEFAULT (0),
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `FK_Community_Creator` (`creator_id`),
  CONSTRAINT `FK_Community_Creator` FOREIGN KEY (`creator_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `community`
--

LOCK TABLES `community` WRITE;
/*!40000 ALTER TABLE `community` DISABLE KEYS */;
/*!40000 ALTER TABLE `community` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `communityinvite`
--

DROP TABLE IF EXISTS `communityinvite`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `communityinvite` (
  `id` int NOT NULL AUTO_INCREMENT,
  `creator_id` int NOT NULL,
  `invited_user_id` int NOT NULL,
  `community_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_communityinvite_creator` (`creator_id`),
  KEY `fk_communityinvite_invited_user` (`invited_user_id`),
  KEY `fk_communityinvite_community` (`community_id`),
  CONSTRAINT `fk_communityinvite_community` FOREIGN KEY (`community_id`) REFERENCES `community` (`id`),
  CONSTRAINT `fk_communityinvite_creator` FOREIGN KEY (`creator_id`) REFERENCES `user` (`id`),
  CONSTRAINT `fk_communityinvite_invited_user` FOREIGN KEY (`invited_user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `communityinvite`
--

LOCK TABLES `communityinvite` WRITE;
/*!40000 ALTER TABLE `communityinvite` DISABLE KEYS */;
/*!40000 ALTER TABLE `communityinvite` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `communitymembership`
--

DROP TABLE IF EXISTS `communitymembership`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `communitymembership` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `community_id` int NOT NULL,
  `is_administrator` tinyint(1) NOT NULL DEFAULT '0',
  `is_moderator` tinyint(1) NOT NULL DEFAULT '0',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_communitymembership_user` (`user_id`),
  KEY `fk_communitymembership_community` (`community_id`),
  CONSTRAINT `fk_communitymembership_community` FOREIGN KEY (`community_id`) REFERENCES `community` (`id`),
  CONSTRAINT `fk_communitymembership_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `communitymembership`
--

LOCK TABLES `communitymembership` WRITE;
/*!40000 ALTER TABLE `communitymembership` DISABLE KEYS */;
/*!40000 ALTER TABLE `communitymembership` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `communitynotificationssubscription`
--

DROP TABLE IF EXISTS `communitynotificationssubscription`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `communitynotificationssubscription` (
  `id` int NOT NULL AUTO_INCREMENT,
  `subscriber_id` int NOT NULL,
  `community_id` int NOT NULL,
  `new_post_notifications` bit(1) NOT NULL DEFAULT (0),
  PRIMARY KEY (`id`),
  KEY `fk_communitynotificationssubscription_subscriber` (`subscriber_id`),
  KEY `fk_communitynotificationssubscription_community` (`community_id`),
  CONSTRAINT `fk_communitynotificationssubscription_community` FOREIGN KEY (`community_id`) REFERENCES `community` (`id`),
  CONSTRAINT `fk_communitynotificationssubscription_subscriber` FOREIGN KEY (`subscriber_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `communitynotificationssubscription`
--

LOCK TABLES `communitynotificationssubscription` WRITE;
/*!40000 ALTER TABLE `communitynotificationssubscription` DISABLE KEYS */;
/*!40000 ALTER TABLE `communitynotificationssubscription` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `connection`
--

DROP TABLE IF EXISTS `connection`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `connection` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `target_user_id` int NOT NULL,
  `target_connection_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_connection_user` (`user_id`),
  KEY `fk_connection_target_user` (`target_user_id`),
  KEY `fk_connection_target_connection` (`target_connection_id`),
  CONSTRAINT `fk_connection_target_connection` FOREIGN KEY (`target_connection_id`) REFERENCES `connection` (`id`),
  CONSTRAINT `fk_connection_target_user` FOREIGN KEY (`target_user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `fk_connection_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `connection`
--

LOCK TABLES `connection` WRITE;
/*!40000 ALTER TABLE `connection` DISABLE KEYS */;
/*!40000 ALTER TABLE `connection` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `connectioncircle`
--

DROP TABLE IF EXISTS `connectioncircle`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `connectioncircle` (
  `id` int NOT NULL AUTO_INCREMENT,
  `connection_id` int NOT NULL,
  `circle_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_connectioncircle_circle` (`circle_id`),
  KEY `fk_connectioncircle_connection` (`connection_id`),
  CONSTRAINT `fk_connectioncircle_circle` FOREIGN KEY (`circle_id`) REFERENCES `circle` (`id`),
  CONSTRAINT `fk_connectioncircle_connection` FOREIGN KEY (`connection_id`) REFERENCES `connection` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `connectioncircle`
--

LOCK TABLES `connectioncircle` WRITE;
/*!40000 ALTER TABLE `connectioncircle` DISABLE KEYS */;
/*!40000 ALTER TABLE `connectioncircle` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `device`
--

DROP TABLE IF EXISTS `device`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `device` (
  `id` int NOT NULL AUTO_INCREMENT,
  `owner_id` int NOT NULL,
  `uuid` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `FK_Device_Owner` (`owner_id`),
  CONSTRAINT `FK_Device_Owner` FOREIGN KEY (`owner_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=203 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `device`
--

LOCK TABLES `device` WRITE;
/*!40000 ALTER TABLE `device` DISABLE KEYS */;
INSERT INTO `device` VALUES (201,210,'a15da083ffd8910295b9efd994b3585d43f595835bc18839c1acc358c4bace52','Android SDK built for x86','2020-07-13 16:22:47'),(202,210,'a5ae84a263dcea9503c39c1fe6927b9f032ba2f37f68baa04fa33f0888791fca','sdk_gphone_x86_arm','0000-00-00 00:00:00');
/*!40000 ALTER TABLE `device` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `emoji`
--

DROP TABLE IF EXISTS `emoji`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `emoji` (
  `id` int NOT NULL AUTO_INCREMENT,
  `keyword` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `color` varchar(7) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `image` varbinary(1000) DEFAULT NULL,
  `order` int NOT NULL,
  `group_Id` int DEFAULT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_emoji` (`keyword`),
  KEY `fk_emoji_groupemoji` (`group_Id`),
  CONSTRAINT `fk_emoji_groupemoji` FOREIGN KEY (`group_Id`) REFERENCES `emojigroup` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `emoji`
--

LOCK TABLES `emoji` WRITE;
/*!40000 ALTER TABLE `emoji` DISABLE KEYS */;
/*!40000 ALTER TABLE `emoji` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `emojigroup`
--

DROP TABLE IF EXISTS `emojigroup`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `emojigroup` (
  `id` int NOT NULL AUTO_INCREMENT,
  `keyword` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `color` varchar(7) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `order` int NOT NULL,
  `is_reaction_group` tinyint(1) NOT NULL DEFAULT '0',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_emojigroup` (`keyword`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `emojigroup`
--

LOCK TABLES `emojigroup` WRITE;
/*!40000 ALTER TABLE `emojigroup` DISABLE KEYS */;
/*!40000 ALTER TABLE `emojigroup` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `follow`
--

DROP TABLE IF EXISTS `follow`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `follow` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `followed_user_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_follow_user` (`user_id`),
  KEY `fk_follow_followed_user` (`followed_user_id`),
  CONSTRAINT `fk_follow_followed_user` FOREIGN KEY (`followed_user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `fk_follow_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `follow`
--

LOCK TABLES `follow` WRITE;
/*!40000 ALTER TABLE `follow` DISABLE KEYS */;
/*!40000 ALTER TABLE `follow` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `followrequest`
--

DROP TABLE IF EXISTS `followrequest`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `followrequest` (
  `id` int NOT NULL AUTO_INCREMENT,
  `creator_id` int NOT NULL,
  `target_user_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_followrequest_creator` (`creator_id`),
  KEY `fk_followrequest_target_user` (`target_user_id`),
  CONSTRAINT `fk_followrequest_creator` FOREIGN KEY (`creator_id`) REFERENCES `user` (`id`),
  CONSTRAINT `fk_followrequest_target_user` FOREIGN KEY (`target_user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `followrequest`
--

LOCK TABLES `followrequest` WRITE;
/*!40000 ALTER TABLE `followrequest` DISABLE KEYS */;
/*!40000 ALTER TABLE `followrequest` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `hashtag`
--

DROP TABLE IF EXISTS `hashtag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `hashtag` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `color` varchar(7) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `text_color` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `width` int DEFAULT NULL,
  `height` int DEFAULT NULL,
  `image` varbinary(1000) DEFAULT NULL,
  `emoji_id` int DEFAULT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_hashtag_emoji` (`emoji_id`),
  CONSTRAINT `fk_hashtag_emoji` FOREIGN KEY (`emoji_id`) REFERENCES `emoji` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `hashtag`
--

LOCK TABLES `hashtag` WRITE;
/*!40000 ALTER TABLE `hashtag` DISABLE KEYS */;
/*!40000 ALTER TABLE `hashtag` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `hashtag_post`
--

DROP TABLE IF EXISTS `hashtag_post`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `hashtag_post` (
  `id` int NOT NULL AUTO_INCREMENT,
  `hashtag_id` int NOT NULL,
  `post_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_hashtag_post_hashtag` (`hashtag_id`),
  KEY `fk_hashtag_post_post` (`post_id`),
  CONSTRAINT `fk_hashtag_post_hashtag` FOREIGN KEY (`hashtag_id`) REFERENCES `hashtag` (`id`),
  CONSTRAINT `fk_hashtag_post_post` FOREIGN KEY (`post_id`) REFERENCES `post` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `hashtag_post`
--

LOCK TABLES `hashtag_post` WRITE;
/*!40000 ALTER TABLE `hashtag_post` DISABLE KEYS */;
/*!40000 ALTER TABLE `hashtag_post` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `hashtag_postcomment`
--

DROP TABLE IF EXISTS `hashtag_postcomment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `hashtag_postcomment` (
  `id` int NOT NULL AUTO_INCREMENT,
  `hashtag_id` int NOT NULL,
  `postcomment_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_hashtag_postcomment_hashtag` (`hashtag_id`),
  KEY `fk_hashtag_postcomment_postcomment` (`postcomment_id`),
  CONSTRAINT `fk_hashtag_postcomment_hashtag` FOREIGN KEY (`hashtag_id`) REFERENCES `hashtag` (`id`),
  CONSTRAINT `fk_hashtag_postcomment_postcomment` FOREIGN KEY (`postcomment_id`) REFERENCES `postcomment` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `hashtag_postcomment`
--

LOCK TABLES `hashtag_postcomment` WRITE;
/*!40000 ALTER TABLE `hashtag_postcomment` DISABLE KEYS */;
/*!40000 ALTER TABLE `hashtag_postcomment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `language`
--

DROP TABLE IF EXISTS `language`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `language` (
  `id` int NOT NULL AUTO_INCREMENT,
  `code` varchar(12) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_language` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `language`
--

LOCK TABLES `language` WRITE;
/*!40000 ALTER TABLE `language` DISABLE KEYS */;
INSERT INTO `language` VALUES (1,'en','English','2020-07-11 02:31:46');
/*!40000 ALTER TABLE `language` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `moderatedobject`
--

DROP TABLE IF EXISTS `moderatedobject`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `moderatedobject` (
  `id` int NOT NULL AUTO_INCREMENT,
  `community_id` int NOT NULL,
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `verified` tinyint(1) NOT NULL DEFAULT '0',
  `status` varchar(5) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `category_id` int DEFAULT NULL,
  `content_type` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `object_id` int DEFAULT NULL,
  `content_object` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_moderatedobject_community` (`community_id`),
  KEY `fk_moderatedobject_category` (`category_id`),
  CONSTRAINT `fk_moderatedobject_category` FOREIGN KEY (`category_id`) REFERENCES `moderationcategory` (`id`),
  CONSTRAINT `fk_moderatedobject_community` FOREIGN KEY (`community_id`) REFERENCES `community` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `moderatedobject`
--

LOCK TABLES `moderatedobject` WRITE;
/*!40000 ALTER TABLE `moderatedobject` DISABLE KEYS */;
/*!40000 ALTER TABLE `moderatedobject` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `moderationcategory`
--

DROP TABLE IF EXISTS `moderationcategory`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `moderationcategory` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `title` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `severity` varchar(5) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `Order` int DEFAULT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `moderationcategory`
--

LOCK TABLES `moderationcategory` WRITE;
/*!40000 ALTER TABLE `moderationcategory` DISABLE KEYS */;
/*!40000 ALTER TABLE `moderationcategory` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `moderationpenalty`
--

DROP TABLE IF EXISTS `moderationpenalty`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `moderationpenalty` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `expiration` timestamp NOT NULL,
  `moderated_object_id` int NOT NULL,
  `type` varchar(5) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_moderationpenalty_user` (`user_id`),
  KEY `fk_moderationpenalty_moderatedobject` (`moderated_object_id`),
  CONSTRAINT `fk_moderationpenalty_moderatedobject` FOREIGN KEY (`moderated_object_id`) REFERENCES `moderatedobject` (`id`),
  CONSTRAINT `fk_moderationpenalty_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `moderationpenalty`
--

LOCK TABLES `moderationpenalty` WRITE;
/*!40000 ALTER TABLE `moderationpenalty` DISABLE KEYS */;
/*!40000 ALTER TABLE `moderationpenalty` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `moderationreport`
--

DROP TABLE IF EXISTS `moderationreport`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `moderationreport` (
  `id` int NOT NULL AUTO_INCREMENT,
  `reporter_id` int NOT NULL,
  `moderated_object_id` int NOT NULL,
  `category_id` int DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_moderationreport_reporter` (`reporter_id`),
  KEY `fk_moderationreport_moderatedobject` (`moderated_object_id`),
  KEY `FK_ModerationReport_Category` (`category_id`),
  CONSTRAINT `FK_ModerationReport_Category` FOREIGN KEY (`category_id`) REFERENCES `moderationcategory` (`id`),
  CONSTRAINT `fk_moderationreport_moderatedobject` FOREIGN KEY (`moderated_object_id`) REFERENCES `moderatedobject` (`id`),
  CONSTRAINT `fk_moderationreport_reporter` FOREIGN KEY (`reporter_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `moderationreport`
--

LOCK TABLES `moderationreport` WRITE;
/*!40000 ALTER TABLE `moderationreport` DISABLE KEYS */;
/*!40000 ALTER TABLE `moderationreport` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `post`
--

DROP TABLE IF EXISTS `post`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `post` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uuid` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `text` varchar(560) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `modified` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `creator_id` int NOT NULL,
  `comments_enabled` tinyint(1) NOT NULL DEFAULT '0',
  `public_reactions` tinyint(1) NOT NULL DEFAULT '0',
  `community_id` int DEFAULT NULL,
  `language_id` int DEFAULT NULL,
  `is_edited` tinyint(1) NOT NULL DEFAULT '0',
  `is_closed` tinyint(1) NOT NULL DEFAULT '0',
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  `status` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `media_height` int DEFAULT NULL,
  `media_width` int DEFAULT NULL,
  `media_thumbnail` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_post_creator` (`creator_id`),
  KEY `fk_post_community` (`community_id`),
  KEY `fk_post_language` (`language_id`),
  CONSTRAINT `fk_post_community` FOREIGN KEY (`community_id`) REFERENCES `community` (`id`),
  CONSTRAINT `fk_post_creator` FOREIGN KEY (`creator_id`) REFERENCES `user` (`id`),
  CONSTRAINT `fk_post_language` FOREIGN KEY (`language_id`) REFERENCES `language` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=214 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `post`
--

LOCK TABLES `post` WRITE;
/*!40000 ALTER TABLE `post` DISABLE KEYS */;
INSERT INTO `post` VALUES (201,'8f117402-c1d3-42cf-b322-f56859194b9f','This is first post','2020-07-13 16:22:47','2020-07-13 16:22:47',210,0,0,NULL,NULL,0,0,0,'P',0,0,''),(202,'b3c04040-090e-408a-a4da-eae593c38bb0','This is second post','2020-07-13 15:43:24','2020-07-13 15:44:00',210,0,0,NULL,NULL,0,0,0,'D',0,0,''),(203,'04bd4245-b3f4-4563-a855-3371702ca09f','This is third post','2020-07-13 15:44:23','2020-07-13 15:44:23',210,0,0,NULL,NULL,0,0,0,'D',0,0,''),(204,'a81f0193-7ec9-467a-99fd-3db19087456c','Nguyen Do comment','2020-07-19 16:49:25','0000-00-00 00:00:00',210,0,0,NULL,NULL,0,0,0,'P',0,0,''),(205,'1a481fe6-4613-48c9-91ea-a91babd558bc','Do Nguyen','2020-07-19 16:49:29','0000-00-00 00:00:00',210,0,0,NULL,NULL,0,0,0,'P',0,0,''),(206,'b6a3713f-4522-4b21-b400-9e08874b840b','Comment 1','2020-07-19 16:55:55','0000-00-00 00:00:00',210,0,0,NULL,NULL,0,0,0,'P',0,0,''),(207,'f0ea6e40-db3c-4d9f-b941-3e39c1d42202','Comment 2','2020-07-19 17:00:54','0000-00-00 00:00:00',210,0,0,NULL,NULL,0,0,0,'P',0,0,''),(208,'46e9b497-b175-4b34-926b-730fecf111af','test 1111','2020-07-19 17:02:48','0000-00-00 00:00:00',210,0,0,NULL,NULL,0,0,0,'P',0,0,''),(209,'da239ca8-ec60-48a0-acf5-33faab32af2a','Comment 2','2020-07-19 17:02:48','0000-00-00 00:00:00',210,0,0,NULL,NULL,0,0,0,'P',0,0,''),(210,'bc02c2fc-90ff-4d16-9e0e-836b464acf0d','test 2222','2020-07-19 17:05:19','0000-00-00 00:00:00',210,0,0,NULL,NULL,0,0,0,'P',0,0,''),(211,'8ba72f8f-a027-49b4-95fe-0e81bb131809','test 333','2020-07-19 17:08:43','0000-00-00 00:00:00',210,0,0,NULL,NULL,0,0,0,'P',0,0,''),(212,'6344724c-fb4b-4a3a-ac3d-6117c87447bf','comment 444','2020-07-19 17:14:02','0000-00-00 00:00:00',210,0,0,NULL,NULL,0,0,0,'P',0,0,''),(213,'90c144c4-a7d3-4819-a7cc-8b8e1d33f66c','Demo 123','2020-12-20 15:22:38','0000-00-00 00:00:00',210,0,0,NULL,NULL,0,0,0,'P',0,0,'');
/*!40000 ALTER TABLE `post` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `postcomment`
--

DROP TABLE IF EXISTS `postcomment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `postcomment` (
  `id` int NOT NULL AUTO_INCREMENT,
  `post_id` int NOT NULL,
  `parent_comment_id` int DEFAULT NULL,
  `commenter_id` int NOT NULL,
  `text` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `language_id` int DEFAULT NULL,
  `is_edited` tinyint(1) NOT NULL DEFAULT '0',
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `modified` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_postcomment_post` (`post_id`),
  KEY `fk_postcomment_parent_comment` (`parent_comment_id`),
  KEY `fk_postcomment_commenter` (`commenter_id`),
  KEY `fk_postcomment_language` (`language_id`),
  CONSTRAINT `fk_postcomment_commenter` FOREIGN KEY (`commenter_id`) REFERENCES `user` (`id`),
  CONSTRAINT `fk_postcomment_language` FOREIGN KEY (`language_id`) REFERENCES `language` (`id`),
  CONSTRAINT `fk_postcomment_parent_comment` FOREIGN KEY (`parent_comment_id`) REFERENCES `postcomment` (`id`),
  CONSTRAINT `fk_postcomment_post` FOREIGN KEY (`post_id`) REFERENCES `post` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `postcomment`
--

LOCK TABLES `postcomment` WRITE;
/*!40000 ALTER TABLE `postcomment` DISABLE KEYS */;
/*!40000 ALTER TABLE `postcomment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `postcommentmute`
--

DROP TABLE IF EXISTS `postcommentmute`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `postcommentmute` (
  `id` int NOT NULL AUTO_INCREMENT,
  `post_comment_id` int NOT NULL,
  `muter_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_postcommentmute_postcomment` (`post_comment_id`),
  KEY `fk_postcommentmute_muter` (`muter_id`),
  CONSTRAINT `fk_postcommentmute_muter` FOREIGN KEY (`muter_id`) REFERENCES `user` (`id`),
  CONSTRAINT `fk_postcommentmute_postcomment` FOREIGN KEY (`post_comment_id`) REFERENCES `postcomment` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `postcommentmute`
--

LOCK TABLES `postcommentmute` WRITE;
/*!40000 ALTER TABLE `postcommentmute` DISABLE KEYS */;
/*!40000 ALTER TABLE `postcommentmute` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `postcommentreaction`
--

DROP TABLE IF EXISTS `postcommentreaction`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `postcommentreaction` (
  `id` int NOT NULL AUTO_INCREMENT,
  `reactor_id` int NOT NULL,
  `emoji_id` int NOT NULL,
  `post_comment_id` int DEFAULT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_postcommentreaction_commenter` (`reactor_id`),
  KEY `fk_postcommentreaction_emoji` (`emoji_id`),
  KEY `fk_postcommentreaction_postcomment` (`post_comment_id`),
  CONSTRAINT `fk_postcommentreaction_commenter` FOREIGN KEY (`reactor_id`) REFERENCES `user` (`id`),
  CONSTRAINT `fk_postcommentreaction_emoji` FOREIGN KEY (`emoji_id`) REFERENCES `emoji` (`id`),
  CONSTRAINT `fk_postcommentreaction_postcomment` FOREIGN KEY (`post_comment_id`) REFERENCES `postcomment` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `postcommentreaction`
--

LOCK TABLES `postcommentreaction` WRITE;
/*!40000 ALTER TABLE `postcommentreaction` DISABLE KEYS */;
/*!40000 ALTER TABLE `postcommentreaction` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `postcommentusermention`
--

DROP TABLE IF EXISTS `postcommentusermention`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `postcommentusermention` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `post_comment_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_postcommentusermention_user` (`user_id`),
  KEY `fk_postcommentusermention_postcomment` (`post_comment_id`),
  CONSTRAINT `fk_postcommentusermention_postcomment` FOREIGN KEY (`post_comment_id`) REFERENCES `postcomment` (`id`),
  CONSTRAINT `fk_postcommentusermention_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `postcommentusermention`
--

LOCK TABLES `postcommentusermention` WRITE;
/*!40000 ALTER TABLE `postcommentusermention` DISABLE KEYS */;
/*!40000 ALTER TABLE `postcommentusermention` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `postimage`
--

DROP TABLE IF EXISTS `postimage`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `postimage` (
  `id` int NOT NULL AUTO_INCREMENT,
  `post_id` int NOT NULL,
  `image` varbinary(1000) DEFAULT NULL,
  `height` int DEFAULT NULL,
  `width` int DEFAULT NULL,
  `hash` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `thumbnail` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_postimage_post` (`post_id`),
  CONSTRAINT `fk_postimage_post` FOREIGN KEY (`post_id`) REFERENCES `post` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `postimage`
--

LOCK TABLES `postimage` WRITE;
/*!40000 ALTER TABLE `postimage` DISABLE KEYS */;
/*!40000 ALTER TABLE `postimage` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `postmedia`
--

DROP TABLE IF EXISTS `postmedia`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `postmedia` (
  `id` int NOT NULL AUTO_INCREMENT,
  `post_id` int NOT NULL,
  `type` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `content_type` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `object_id` int DEFAULT NULL,
  `content_object` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_postmedia_post` (`post_id`),
  CONSTRAINT `fk_postmedia_post` FOREIGN KEY (`post_id`) REFERENCES `post` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `postmedia`
--

LOCK TABLES `postmedia` WRITE;
/*!40000 ALTER TABLE `postmedia` DISABLE KEYS */;
/*!40000 ALTER TABLE `postmedia` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `postmute`
--

DROP TABLE IF EXISTS `postmute`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `postmute` (
  `id` int NOT NULL AUTO_INCREMENT,
  `post_id` int NOT NULL,
  `muter_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_postmute_post` (`post_id`),
  KEY `fk_postmute_muter` (`muter_id`),
  CONSTRAINT `fk_postmute_muter` FOREIGN KEY (`muter_id`) REFERENCES `user` (`id`),
  CONSTRAINT `fk_postmute_post` FOREIGN KEY (`post_id`) REFERENCES `post` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `postmute`
--

LOCK TABLES `postmute` WRITE;
/*!40000 ALTER TABLE `postmute` DISABLE KEYS */;
/*!40000 ALTER TABLE `postmute` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `postreaction`
--

DROP TABLE IF EXISTS `postreaction`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `postreaction` (
  `id` int NOT NULL AUTO_INCREMENT,
  `post_id` int NOT NULL,
  `reactor_id` int NOT NULL,
  `emoji_id` int NOT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_postreaction_post` (`post_id`),
  KEY `fk_postreaction_commenter` (`reactor_id`),
  KEY `fk_postreaction_emoji` (`emoji_id`),
  CONSTRAINT `fk_postreaction_commenter` FOREIGN KEY (`reactor_id`) REFERENCES `user` (`id`),
  CONSTRAINT `fk_postreaction_emoji` FOREIGN KEY (`emoji_id`) REFERENCES `emoji` (`id`),
  CONSTRAINT `fk_postreaction_post` FOREIGN KEY (`post_id`) REFERENCES `post` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `postreaction`
--

LOCK TABLES `postreaction` WRITE;
/*!40000 ALTER TABLE `postreaction` DISABLE KEYS */;
/*!40000 ALTER TABLE `postreaction` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `postusermention`
--

DROP TABLE IF EXISTS `postusermention`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `postusermention` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `post_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_postusermention_user` (`user_id`),
  KEY `fk_postusermention_post` (`post_id`),
  CONSTRAINT `fk_postusermention_post` FOREIGN KEY (`post_id`) REFERENCES `post` (`id`),
  CONSTRAINT `fk_postusermention_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `postusermention`
--

LOCK TABLES `postusermention` WRITE;
/*!40000 ALTER TABLE `postusermention` DISABLE KEYS */;
/*!40000 ALTER TABLE `postusermention` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `postvideo`
--

DROP TABLE IF EXISTS `postvideo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `postvideo` (
  `id` int NOT NULL AUTO_INCREMENT,
  `post_id` int NOT NULL,
  `hash` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `height` int DEFAULT NULL,
  `width` int DEFAULT NULL,
  `duration` float DEFAULT NULL,
  `file` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `file_format` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `thumbnail` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `thumbnail_height` int DEFAULT NULL,
  `thumbnail_width` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_postvideo_post` (`post_id`),
  CONSTRAINT `fk_postvideo_post` FOREIGN KEY (`post_id`) REFERENCES `post` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `postvideo`
--

LOCK TABLES `postvideo` WRITE;
/*!40000 ALTER TABLE `postvideo` DISABLE KEYS */;
/*!40000 ALTER TABLE `postvideo` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `proxyblacklisteddomain`
--

DROP TABLE IF EXISTS `proxyblacklisteddomain`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `proxyblacklisteddomain` (
  `id` int NOT NULL AUTO_INCREMENT,
  `domain` varchar(150) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `proxyblacklisteddomain`
--

LOCK TABLES `proxyblacklisteddomain` WRITE;
/*!40000 ALTER TABLE `proxyblacklisteddomain` DISABLE KEYS */;
/*!40000 ALTER TABLE `proxyblacklisteddomain` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `toppost`
--

DROP TABLE IF EXISTS `toppost`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `toppost` (
  `id` int NOT NULL AUTO_INCREMENT,
  `post_id` int NOT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_toppost_post` (`post_id`),
  CONSTRAINT `fk_toppost_post` FOREIGN KEY (`post_id`) REFERENCES `post` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `toppost`
--

LOCK TABLES `toppost` WRITE;
/*!40000 ALTER TABLE `toppost` DISABLE KEYS */;
INSERT INTO `toppost` VALUES (1,201,'2020-07-13 15:43:24'),(2,202,'2020-07-13 15:43:30');
/*!40000 ALTER TABLE `toppost` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `toppostcommunityexclusion`
--

DROP TABLE IF EXISTS `toppostcommunityexclusion`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `toppostcommunityexclusion` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `community_id` int DEFAULT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_toppostcommunityexclusion_user` (`user_id`),
  KEY `fk_toppostcommunityexclusion_community` (`community_id`),
  CONSTRAINT `fk_toppostcommunityexclusion_community` FOREIGN KEY (`community_id`) REFERENCES `community` (`id`),
  CONSTRAINT `fk_toppostcommunityexclusion_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `toppostcommunityexclusion`
--

LOCK TABLES `toppostcommunityexclusion` WRITE;
/*!40000 ALTER TABLE `toppostcommunityexclusion` DISABLE KEYS */;
/*!40000 ALTER TABLE `toppostcommunityexclusion` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `trendingpost`
--

DROP TABLE IF EXISTS `trendingpost`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `trendingpost` (
  `id` int NOT NULL AUTO_INCREMENT,
  `post_id` int NOT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_trendingpost_post` (`post_id`),
  CONSTRAINT `fk_trendingpost_post` FOREIGN KEY (`post_id`) REFERENCES `post` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `trendingpost`
--

LOCK TABLES `trendingpost` WRITE;
/*!40000 ALTER TABLE `trendingpost` DISABLE KEYS */;
/*!40000 ALTER TABLE `trendingpost` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `password` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `email` varchar(254) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `last_login` timestamp NOT NULL,
  `is_superuser` tinyint(1) NOT NULL DEFAULT '0',
  `is_staff` tinyint(1) NOT NULL DEFAULT '0',
  `is_active` tinyint(1) NOT NULL DEFAULT '0',
  `date_joined` timestamp NOT NULL,
  `is_email_verified` tinyint(1) NOT NULL DEFAULT '0',
  `are_guidelines_accepted` tinyint(1) NOT NULL DEFAULT '0',
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  `visibility` varchar(2) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `invite_count` int NOT NULL DEFAULT (0),
  `language_id` int DEFAULT NULL,
  `translation_language_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_user_language` (`language_id`),
  KEY `fk_user_translation_language` (`translation_language_id`),
  CONSTRAINT `fk_user_language` FOREIGN KEY (`language_id`) REFERENCES `language` (`id`),
  CONSTRAINT `fk_user_translation_language` FOREIGN KEY (`translation_language_id`) REFERENCES `language` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=211 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (210,'don','$2a$10$ORGoX7okYa2IZpdN9l.k2.TKXATR0GBZE7QNGHfh2VHFtC8FS2Wy6','nguyendo2008@gmail.com','0000-00-00 00:00:00',0,0,0,'0000-00-00 00:00:00',0,1,0,'',0,NULL,NULL);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `userblock`
--

DROP TABLE IF EXISTS `userblock`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `userblock` (
  `id` int NOT NULL AUTO_INCREMENT,
  `blocked_user_id` int NOT NULL,
  `blocker_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_userblock_blockeduser` (`blocked_user_id`),
  KEY `fk_userblock_blocker` (`blocker_id`),
  CONSTRAINT `fk_userblock_blockeduser` FOREIGN KEY (`blocked_user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `fk_userblock_blocker` FOREIGN KEY (`blocker_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `userblock`
--

LOCK TABLES `userblock` WRITE;
/*!40000 ALTER TABLE `userblock` DISABLE KEYS */;
/*!40000 ALTER TABLE `userblock` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `userinvite`
--

DROP TABLE IF EXISTS `userinvite`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `userinvite` (
  `id` int NOT NULL AUTO_INCREMENT,
  `invited_by_id` int NOT NULL,
  `created_user_id` int DEFAULT NULL,
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `nickname` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `email` varchar(254) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `username` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `token` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `badge_id` int DEFAULT NULL,
  `is_invite_email_sent` tinyint(1) NOT NULL DEFAULT '0',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `FK_UserInvite_Invited_By` (`invited_by_id`),
  KEY `FK_UserInvite_Created_User` (`created_user_id`),
  KEY `FK_UserInvite_Badge` (`badge_id`),
  CONSTRAINT `FK_UserInvite_Badge` FOREIGN KEY (`badge_id`) REFERENCES `badge` (`id`),
  CONSTRAINT `FK_UserInvite_Created_User` FOREIGN KEY (`created_user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `FK_UserInvite_Invited_By` FOREIGN KEY (`invited_by_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `userinvite`
--

LOCK TABLES `userinvite` WRITE;
/*!40000 ALTER TABLE `userinvite` DISABLE KEYS */;
/*!40000 ALTER TABLE `userinvite` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `usernotificationssettings`
--

DROP TABLE IF EXISTS `usernotificationssettings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `usernotificationssettings` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `post_comment_notifications` tinyint(1) NOT NULL DEFAULT '0',
  `post_comment_reply_notifications` tinyint(1) NOT NULL DEFAULT '0',
  `post_reaction_notifications` tinyint(1) NOT NULL DEFAULT '0',
  `follow_notifications` tinyint(1) NOT NULL DEFAULT '0',
  `follow_request_notifications` tinyint(1) NOT NULL DEFAULT '0',
  `follow_request_approved_notifications` tinyint(1) NOT NULL DEFAULT '0',
  `connection_request_notifications` tinyint(1) NOT NULL DEFAULT '0',
  `connection_confirmed_notifications` tinyint(1) NOT NULL DEFAULT '0',
  `community_invite_notifications` tinyint(1) NOT NULL DEFAULT '0',
  `community_new_post_notifications` tinyint(1) NOT NULL DEFAULT '0',
  `user_new_post_notifications` tinyint(1) NOT NULL DEFAULT '0',
  `post_comment_reaction_notifications` tinyint(1) NOT NULL DEFAULT '0',
  `post_comment_user_mention_notifications` tinyint(1) NOT NULL DEFAULT '0',
  `post_User_Mention_Notifications` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `fk_usernotificationssettings_user` (`user_id`),
  CONSTRAINT `fk_usernotificationssettings_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `usernotificationssettings`
--

LOCK TABLES `usernotificationssettings` WRITE;
/*!40000 ALTER TABLE `usernotificationssettings` DISABLE KEYS */;
/*!40000 ALTER TABLE `usernotificationssettings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `usernotificationssubscription`
--

DROP TABLE IF EXISTS `usernotificationssubscription`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `usernotificationssubscription` (
  `id` int NOT NULL AUTO_INCREMENT,
  `subscriber_id` int NOT NULL,
  `user_id` int NOT NULL,
  `new_post_notifications` bit(1) NOT NULL DEFAULT (0),
  PRIMARY KEY (`id`),
  KEY `fk_usernotificationssubscription_subscriber` (`subscriber_id`),
  KEY `fk_usernotificationssubscription_user` (`user_id`),
  CONSTRAINT `fk_usernotificationssubscription_subscriber` FOREIGN KEY (`subscriber_id`) REFERENCES `user` (`id`),
  CONSTRAINT `fk_usernotificationssubscription_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `usernotificationssubscription`
--

LOCK TABLES `usernotificationssubscription` WRITE;
/*!40000 ALTER TABLE `usernotificationssubscription` DISABLE KEYS */;
/*!40000 ALTER TABLE `usernotificationssubscription` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `userprofile`
--

DROP TABLE IF EXISTS `userprofile`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `userprofile` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `name` varchar(192) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `location` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `is_of_legal_age` tinyint(1) NOT NULL DEFAULT '0',
  `avatar` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `cover` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `bio` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `url` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `followers_count_visible` tinyint(1) NOT NULL DEFAULT '0',
  `community_posts_visible` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  KEY `fk_userprofile_user` (`user_id`),
  CONSTRAINT `fk_userprofile_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `userprofile`
--

LOCK TABLES `userprofile` WRITE;
/*!40000 ALTER TABLE `userprofile` DISABLE KEYS */;
INSERT INTO `userprofile` VALUES (5,210,'Nguyen Do','',1,'https://thirsty-wescoff-17b34d.netlify.app/photo.png','','','',0,0);
/*!40000 ALTER TABLE `userprofile` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `userprofile_badge`
--

DROP TABLE IF EXISTS `userprofile_badge`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `userprofile_badge` (
  `id` int NOT NULL AUTO_INCREMENT,
  `userprofile_id` int NOT NULL,
  `badge_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_userprofile_badge_userprofile` (`userprofile_id`),
  KEY `fk_userprofile_badge_badge` (`badge_id`),
  CONSTRAINT `fk_userprofile_badge_badge` FOREIGN KEY (`badge_id`) REFERENCES `badge` (`id`),
  CONSTRAINT `fk_userprofile_badge_userprofile` FOREIGN KEY (`userprofile_id`) REFERENCES `userprofile` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `userprofile_badge`
--

LOCK TABLES `userprofile_badge` WRITE;
/*!40000 ALTER TABLE `userprofile_badge` DISABLE KEYS */;
/*!40000 ALTER TABLE `userprofile_badge` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'hummingbird'
--

--
-- Dumping routines for database 'hummingbird'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-01-03 11:31:20
