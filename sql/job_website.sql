-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jul 30, 2023 at 04:34 PM
-- Server version: 10.4.24-MariaDB
-- PHP Version: 8.1.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `job_website`
--

-- --------------------------------------------------------

--
-- Table structure for table `apply_lamaran_users`
--

CREATE TABLE `apply_lamaran_users` (
  `id_user` int(11) NOT NULL,
  `pesan` text NOT NULL,
  `id_company` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `benefit_lowongan_perusahaans`
--

CREATE TABLE `benefit_lowongan_perusahaans` (
  `id_company` int(11) NOT NULL,
  `nama` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `detail_perusahaans`
--

CREATE TABLE `detail_perusahaans` (
  `id_company` int(11) NOT NULL,
  `nama` varchar(255) DEFAULT NULL,
  `alamat` text DEFAULT NULL,
  `deskripsi` text DEFAULT NULL,
  `bidang` text DEFAULT NULL,
  `pencapaian` text DEFAULT NULL,
  `jumlah_karyawan` bigint(20) DEFAULT NULL,
  `website` varchar(255) DEFAULT NULL,
  `logo` varchar(255) DEFAULT NULL,
  `background` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `detail_users`
--

CREATE TABLE `detail_users` (
  `id_user` int(11) NOT NULL,
  `gender` varchar(255) DEFAULT NULL,
  `usia` bigint(20) DEFAULT NULL,
  `no_hp` varchar(255) DEFAULT NULL,
  `alamat` text DEFAULT NULL,
  `tanggal_lahir` date DEFAULT NULL,
  `cv` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `keahlian_users`
--

CREATE TABLE `keahlian_users` (
  `id_user` int(11) NOT NULL,
  `nama_keahlian` varchar(255) NOT NULL,
  `level` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `lowongan_perusahaans`
--

CREATE TABLE `lowongan_perusahaans` (
  `id_company` int(11) NOT NULL,
  `title` varchar(255) NOT NULL,
  `deskripsi` text NOT NULL,
  `min_gaji` int(11) NOT NULL,
  `max_gaji` int(11) NOT NULL,
  `poster` varchar(255) NOT NULL,
  `durasi_lowongan` date NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `pendidikan_formal_users`
--

CREATE TABLE `pendidikan_formal_users` (
  `id_user` int(11) NOT NULL,
  `nama_sekolah` varchar(255) DEFAULT NULL,
  `tanggal_masuk` date DEFAULT NULL,
  `tanggal_lulus` date DEFAULT NULL,
  `tingkat_pendidikan` varchar(255) DEFAULT NULL,
  `jurusan` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `pendidikan_non_formal_users`
--

CREATE TABLE `pendidikan_non_formal_users` (
  `id_user` int(11) NOT NULL,
  `penyelenggara` varchar(255) DEFAULT NULL,
  `tanggal_masuk` date DEFAULT NULL,
  `tanggal_lulus` date DEFAULT NULL,
  `jurusan` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `pengalaman_users`
--

CREATE TABLE `pengalaman_users` (
  `id_user` int(11) DEFAULT NULL,
  `nama_perusahaan` varchar(255) DEFAULT NULL,
  `tanggal_masuk` date DEFAULT NULL,
  `tanggal_keluar` date DEFAULT NULL,
  `posisi_terakhir` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `perusahaans`
--

CREATE TABLE `perusahaans` (
  `id_company` int(11) NOT NULL,
  `nama` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `role` varchar(255) DEFAULT 'company',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `requirement_lowongan_perusahaans`
--

CREATE TABLE `requirement_lowongan_perusahaans` (
  `id_company` int(11) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `deskripsi` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id_user` int(11) NOT NULL,
  `first_name` varchar(255) DEFAULT NULL,
  `last_name` varchar(255) DEFAULT NULL,
  `username` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `specialist` varchar(255) DEFAULT NULL,
  `role` varchar(255) DEFAULT 'user',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id_user`, `first_name`, `last_name`, `username`, `email`, `password`, `specialist`, `role`, `created_at`, `updated_at`) VALUES
(2, 'heii', 'test', 'test', 'test1@mail.com', '$argon2id$v=19$m=65536,t=3,p=2$14axJPPsJfJC0gPIAsDTgw$0qnzngAqPsKJcRdgJsIMr+HAW/pL4P987Uek7U136cA', 'test', 'user', '2023-07-29 08:22:54', '2023-07-29 08:22:54'),
(3, 'heii', 'test', 'test', 'test12@mail.com', '$argon2id$v=19$m=65536,t=3,p=2$KNZuLE4JUIHEYQt3w4wfQg$GGhyLNab6hGU/miI+psp4R0XKPe9dLjJmuk7GfbKQJc', 'test', 'user', '2023-07-29 08:23:05', '2023-07-29 08:23:05');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `perusahaans`
--
ALTER TABLE `perusahaans`
  ADD PRIMARY KEY (`id_company`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id_user`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `perusahaans`
--
ALTER TABLE `perusahaans`
  MODIFY `id_company` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id_user` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;