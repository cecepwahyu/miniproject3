-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jun 09, 2023 at 11:28 AM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `miniproject2`
--

-- --------------------------------------------------------

--
-- Table structure for table `account`
--

CREATE TABLE `account` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `role_id` int(10) UNSIGNED DEFAULT NULL,
  `verified` enum('true','false') NOT NULL DEFAULT 'false',
  `active` enum('true','false') NOT NULL DEFAULT 'false',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `account`
--

INSERT INTO `account` (`id`, `username`, `password`, `role_id`, `verified`, `active`, `created_at`, `updated_at`) VALUES
(3, 'test1', 'test1234', 2, 'true', 'true', '2023-06-07 04:30:10', '2023-06-07 04:30:10'),
(4, 'ub_ahz', 'test1234', 2, 'false', 'false', '2023-06-07 05:19:33', '2023-06-07 15:47:27'),
(5, 'dockerfile', '$2a$10$ZpsGKefVXIX5g.hOJBSftu1GifOVgOMxN9S22vXhPxH9XmzAfKtIO', 2, '', 'true', '2023-06-07 04:59:48', '2023-06-07 15:29:17'),
(6, '', '$2a$10$99lQ653CyAD5Lfho6H08ue6D4CmNfsqshBXkUOwI33kqzyGA8PokW', 2, '', 'false', '2023-06-07 05:12:09', '2023-06-07 15:36:31'),
(7, 'dockerfile', '$2a$10$rmTOXi36sVXxMp8388kpFOPg9ibCpIeeja6dIkc1GGza7gwgy4Rq.', 2, '', 'false', '2023-06-07 04:58:35', '2023-06-07 15:23:14');

-- --------------------------------------------------------

--
-- Table structure for table `customer`
--

CREATE TABLE `customer` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `first_name` varchar(255) DEFAULT NULL,
  `last_name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `customer`
--

INSERT INTO `customer` (`id`, `first_name`, `last_name`, `email`, `avatar`, `created_at`, `updated_at`) VALUES
(1, 'jhon', 'jhon', 'email', 'test123', '2023-06-04 21:18:49', '2023-06-04 21:18:49'),
(2, 'jhon', 'jhon', 'email', 'test123', '2023-06-05 03:40:49', '2023-06-05 03:40:49'),
(3, 'jhon', 'jhon', 'email', 'test123', '2023-06-05 05:54:51', '2023-06-05 05:54:51'),
(4, 'jhon', 'jhon', 'email', 'test123', '2023-06-05 23:27:06', '2023-06-05 23:27:06'),
(5, 'jhon', 'jhon', 'email', 'test123', '2023-06-05 23:27:32', '2023-06-05 23:27:32'),
(6, 'jhon', 'jhon', 'email', 'test123', '2023-06-05 23:28:39', '2023-06-05 23:28:39'),
(7, 'jhon', 'jhon', 'email', 'test123', '2023-06-05 23:29:21', '2023-06-05 23:29:21'),
(8, 'jhon', 'jhon', 'email', 'test123', '2023-06-05 23:31:56', '2023-06-05 23:31:56'),
(9, 'jhon', 'jhon', 'email', 'test123', '2023-06-05 23:32:58', '2023-06-05 23:32:58'),
(10, 'jhon', 'cin', 'email', 'test123', '2023-06-05 23:33:27', '2023-06-05 23:33:27');

-- --------------------------------------------------------

--
-- Table structure for table `register_approval`
--

CREATE TABLE `register_approval` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `admin_id` bigint(20) UNSIGNED DEFAULT NULL,
  `super_admin_id` bigint(20) UNSIGNED DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `roles`
--

CREATE TABLE `roles` (
  `id` int(10) UNSIGNED NOT NULL,
  `role_name` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `roles`
--

INSERT INTO `roles` (`id`, `role_name`) VALUES
(1, 'superadmin'),
(2, 'admin');

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `id` int(11) UNSIGNED NOT NULL,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`id`, `name`, `email`, `password`, `created_at`, `updated_at`) VALUES
(1, '', 'email', '', '2023-06-05 08:30:16', '2023-06-05 08:30:16'),
(2, 'jhon', 'email', 'test123', '2023-06-05 08:31:00', '2023-06-05 08:31:00'),
(3, 'ton', 'email', 'test123', '2023-06-05 08:31:36', '2023-06-05 08:31:36'),
(4, '', 'email', 'test123', '2023-06-05 08:48:46', '2023-06-05 08:48:46'),
(5, '', 'email', 'test123', '2023-06-05 17:43:04', '2023-06-05 17:43:04'),
(6, '', 'email', 'test123', '2023-06-05 22:06:40', '2023-06-05 22:06:40'),
(7, '', '', 'docker', '2023-06-06 21:20:09', '2023-06-06 21:20:09'),
(8, '', '', 'docker', '2023-06-06 21:21:04', '2023-06-06 21:21:04'),
(9, '', '', 'docker', '2023-06-06 21:22:50', '2023-06-06 21:22:50'),
(10, '', '', 'docker', '2023-06-06 21:22:55', '2023-06-06 21:22:55'),
(11, '', '', 'docker', '2023-06-06 21:24:07', '2023-06-06 21:24:07'),
(12, 'dockerfile', 'docker', '123', '2023-06-07 04:57:34', '2023-06-07 04:57:34');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `account`
--
ALTER TABLE `account`
  ADD PRIMARY KEY (`id`),
  ADD KEY `role_id` (`role_id`);

--
-- Indexes for table `customer`
--
ALTER TABLE `customer`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `register_approval`
--
ALTER TABLE `register_approval`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_admin` (`admin_id`),
  ADD KEY `fk_superadmin` (`super_admin_id`);

--
-- Indexes for table `roles`
--
ALTER TABLE `roles`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `account`
--
ALTER TABLE `account`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `customer`
--
ALTER TABLE `customer`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `account`
--
ALTER TABLE `account`
  ADD CONSTRAINT `account_ibfk_1` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`),
  ADD CONSTRAINT `account_ibfk_2` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);

--
-- Constraints for table `register_approval`
--
ALTER TABLE `register_approval`
  ADD CONSTRAINT `fk_admin` FOREIGN KEY (`admin_id`) REFERENCES `account` (`id`),
  ADD CONSTRAINT `fk_superadmin` FOREIGN KEY (`super_admin_id`) REFERENCES `account` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
