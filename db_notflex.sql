-- phpMyAdmin SQL Dump
-- version 5.0.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Apr 27, 2021 at 03:44 PM
-- Server version: 10.4.11-MariaDB
-- PHP Version: 7.4.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_notflex`
--

-- --------------------------------------------------------

--
-- Table structure for table `film`
--

CREATE TABLE `film` (
  `id_film` int(5) NOT NULL,
  `judul` varchar(100) NOT NULL,
  `tahun` int(11) NOT NULL,
  `genre` varchar(100) NOT NULL,
  `sutradara` varchar(100) NOT NULL,
  `pemain_utama` varchar(100) NOT NULL,
  `sinopsis` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `film`
--

INSERT INTO `film` (`id_film`, `judul`, `tahun`, `genre`, `sutradara`, `pemain_utama`, `sinopsis`) VALUES
(1, 'Phantom Blood', 2011, 'Action', 'Hirohiko Araki', 'Jonathan Joestar', 'Joestar memburu vampir bernama Dio.'),
(2, 'Battle Tendency', 2013, 'Comedy', 'Hirohiko Araki', 'Joseph Joestar', 'Cucu Jonathan memburu manusia pilar.'),
(3, 'One Piece Whole Series', 2021, 'Adventure', 'Eichiro Oda', 'Monkey D. Luffy', 'Bajak laut tangannnya bisa melar.'),
(4, 'Crazy Rich Rancaekek', 2021, 'Comedy', 'Andreas Virgo', 'Vincent Kurniadi', 'Seorang pemuda Rancaekek jadi kaya.'),
(5, 'Titan Slayer', 2021, 'Horror', 'Komurasaki Daguva', 'Eren Ackerman', 'Seorang pemuda terpaksa harus membunuh Titan.');

-- --------------------------------------------------------

--
-- Table structure for table `history`
--

CREATE TABLE `history` (
  `email_member` varchar(100) NOT NULL,
  `id_film` int(11) NOT NULL,
  `tgl_menonton` date NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `history`
--

INSERT INTO `history` (`email_member`, `id_film`, `tgl_menonton`) VALUES
('alexander@gmail.com', 3, '2021-04-27'),
('aristo@gmail.com', 1, '2021-04-27'),
('aristo@gmail.com', 2, '2021-04-27'),
('aristo@gmail.com', 5, '2021-04-27');

-- --------------------------------------------------------

--
-- Table structure for table `subscription`
--

CREATE TABLE `subscription` (
  `id_langganan` int(5) NOT NULL,
  `email_member` varchar(100) NOT NULL,
  `paket` varchar(100) NOT NULL,
  `no_cc` int(16) NOT NULL,
  `masa_berlaku` varchar(7) NOT NULL,
  `kode_cvc` int(3) NOT NULL,
  `tgl_langganan` date NOT NULL,
  `tgl_berhenti` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `subscription`
--

INSERT INTO `subscription` (`id_langganan`, `email_member`, `paket`, `no_cc`, `masa_berlaku`, `kode_cvc`, `tgl_langganan`, `tgl_berhenti`) VALUES
(1, 'alexander@gmail.com', 'Basic', 1231230151, '04-2024', 337, '2021-04-27', '2021-04-27'),
(2, 'aristo@gmail.com', 'Premium', 2147483647, '04-2024', 662, '2021-04-27', '2021-04-27'),
(3, 'aristo@gmail.com', 'Basic', 2147483647, '04-2024', 212, '2021-04-27', '0000-00-00');

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `email` varchar(100) NOT NULL,
  `nama` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `tgl_lahir` date NOT NULL,
  `jns_kelamin` varchar(100) NOT NULL,
  `asal_negara` varchar(100) NOT NULL,
  `status` varchar(100) NOT NULL,
  `tipe_user` int(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`email`, `nama`, `password`, `tgl_lahir`, `jns_kelamin`, `asal_negara`, `status`, `tipe_user`) VALUES
('alexander@gmail.com', 'Alexander Hansen', 'alex1123', '2000-02-26', 'Pria', 'Indonesia', 'Aktif', 0),
('andreas@gmail.com', 'Andreas V', 'notflex_adm2', '2001-04-07', 'Pria', 'Indonesia', 'Admin', 1),
('aristo@gmail.com', 'Aristo Demos K.', 'thesky', '2000-11-11', 'Pria', 'Myanmar', 'Aktif', 0),
('nivelmart12@gmail.com', 'Levin', 'notflex_adm1', '2001-03-22', 'Pria', 'Indonesia', 'Admin', 1),
('ojan@gmail.com', 'Azareel Fausan', 'ojanjan', '2000-11-11', 'Pria', 'Kamboja', 'Terkunci', 0);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `film`
--
ALTER TABLE `film`
  ADD PRIMARY KEY (`id_film`);

--
-- Indexes for table `history`
--
ALTER TABLE `history`
  ADD PRIMARY KEY (`email_member`,`id_film`),
  ADD KEY `id_film` (`id_film`);

--
-- Indexes for table `subscription`
--
ALTER TABLE `subscription`
  ADD PRIMARY KEY (`id_langganan`),
  ADD KEY `email_member` (`email_member`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`email`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `film`
--
ALTER TABLE `film`
  MODIFY `id_film` int(5) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `subscription`
--
ALTER TABLE `subscription`
  MODIFY `id_langganan` int(5) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `history`
--
ALTER TABLE `history`
  ADD CONSTRAINT `history_ibfk_1` FOREIGN KEY (`email_member`) REFERENCES `user` (`email`),
  ADD CONSTRAINT `history_ibfk_2` FOREIGN KEY (`id_film`) REFERENCES `film` (`id_film`);

--
-- Constraints for table `subscription`
--
ALTER TABLE `subscription`
  ADD CONSTRAINT `subscription_ibfk_1` FOREIGN KEY (`email_member`) REFERENCES `user` (`email`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
