/*
SQLyog Ultimate v13.1.1 (64 bit)
MySQL - 10.4.24-MariaDB : Database - job_website
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`job_website` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `job_website`;

/*Table structure for table `apply_lamaran_users` */

DROP TABLE IF EXISTS `apply_lamaran_users`;

CREATE TABLE `apply_lamaran_users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_user` int(11) NOT NULL,
  `pesan` text NOT NULL,
  `id_company` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;

/*Data for the table `apply_lamaran_users` */

insert  into `apply_lamaran_users`(`id`,`id_user`,`pesan`,`id_company`,`created_at`,`updated_at`) values 
(5,4,'Buka aplikasi Message . Ketuk Tulis Tulis . Di bagian Kepada, masukkan nama, nomor telepon, atau alamat email yang ingin Anda kirimi pesan.',1,'2023-07-31 10:07:11','2023-07-31 10:07:11'),
(7,4,'Buka aplikasi Message . Ketuk Tulis Tulis . Di bagian Kepada, masukkan nama, nomor telepon, atau alamat email yang ingin Anda kirimi pesan.',2,'2023-07-31 10:12:47','2023-07-31 10:12:47');

/*Table structure for table `benefit_lowongan_perusahaans` */

DROP TABLE IF EXISTS `benefit_lowongan_perusahaans`;

CREATE TABLE `benefit_lowongan_perusahaans` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_lowongan` int(11) NOT NULL,
  `nama` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4;

/*Data for the table `benefit_lowongan_perusahaans` */

insert  into `benefit_lowongan_perusahaans`(`id`,`id_lowongan`,`nama`) values 
(7,7,'tangan bersih'),
(8,7,'makan'),
(11,9,'tangan bersih'),
(12,9,'makan'),
(13,10,'tangan bersih'),
(14,10,'makan'),
(15,11,'tangan bersih'),
(16,11,'makan');

/*Table structure for table `detail_perusahaans` */

DROP TABLE IF EXISTS `detail_perusahaans`;

CREATE TABLE `detail_perusahaans` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_company` int(11) NOT NULL,
  `alamat` text DEFAULT NULL,
  `deskripsi` text DEFAULT NULL,
  `bidang` text DEFAULT NULL,
  `pencapaian` text DEFAULT NULL,
  `jumlah_karyawan` bigint(20) DEFAULT NULL,
  `website` varchar(255) DEFAULT NULL,
  `logo` varchar(255) DEFAULT NULL,
  `background` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`,`id_company`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

/*Data for the table `detail_perusahaans` */

insert  into `detail_perusahaans`(`id`,`id_company`,`alamat`,`deskripsi`,`bidang`,`pencapaian`,`jumlah_karyawan`,`website`,`logo`,`background`,`created_at`,`updated_at`) values 
(4,1,'ungsi append() yang memungkinkan Anda menambahkan elemen baru ke dalam slice (slice mirip dengan array, namun ukurannya dapat berubah)','ungsi append() yang memungkinkan Anda menambahkan elemen baru ke dalam slice (slice mirip dengan array, namun ukurannya dapat berubah)','perikanan','',50,'www.google.com','test','test','2023-07-31 14:22:03','2023-07-31 14:22:03');

/*Table structure for table `detail_users` */

DROP TABLE IF EXISTS `detail_users`;

CREATE TABLE `detail_users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_user` int(11) NOT NULL,
  `gender` varchar(255) DEFAULT NULL,
  `usia` bigint(20) DEFAULT NULL,
  `no_hp` varchar(255) DEFAULT NULL,
  `alamat` text DEFAULT NULL,
  `tanggal_lahir` date DEFAULT NULL,
  `cv` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`,`id_user`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

/*Data for the table `detail_users` */

insert  into `detail_users`(`id`,`id_user`,`gender`,`usia`,`no_hp`,`alamat`,`tanggal_lahir`,`cv`,`created_at`,`updated_at`) values 
(1,4,'laki-laki',12,'089456161632','Filipino boy group formed in 2020 by Viva Artist Agency and Ninuno Media. The group consists of six members: Taneo, Mo, Jao, Tomas, R-Ji','2023-07-31','tester','2023-07-31 09:30:10','2023-07-31 09:30:10');

/*Table structure for table `keahlian_users` */

DROP TABLE IF EXISTS `keahlian_users`;

CREATE TABLE `keahlian_users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_user` int(11) NOT NULL,
  `nama_keahlian` varchar(255) NOT NULL,
  `level` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

/*Data for the table `keahlian_users` */

insert  into `keahlian_users`(`id`,`id_user`,`nama_keahlian`,`level`) values 
(1,4,'JAVASCRIPT','EXPERT'),
(2,4,'GOLANG','MEDIUM');

/*Table structure for table `lowongan_perusahaans` */

DROP TABLE IF EXISTS `lowongan_perusahaans`;

CREATE TABLE `lowongan_perusahaans` (
  `id_lowongan` int(11) NOT NULL AUTO_INCREMENT,
  `id_company` int(11) NOT NULL,
  `title` varchar(255) NOT NULL,
  `deskripsi` text NOT NULL,
  `min_gaji` int(11) NOT NULL,
  `max_gaji` int(11) NOT NULL,
  `poster` varchar(255) NOT NULL,
  `durasi_lowongan` date NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id_lowongan`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4;

/*Data for the table `lowongan_perusahaans` */

insert  into `lowongan_perusahaans`(`id_lowongan`,`id_company`,`title`,`deskripsi`,`min_gaji`,`max_gaji`,`poster`,`durasi_lowongan`,`created_at`,`updated_at`) values 
(7,1,'PENCUCI PIRING','mencuci sebagian piring kotor',1000,20000,'test','2023-08-24','2023-07-31 13:47:06','2023-07-31 13:47:06'),
(9,1,'PENCUCI MATA','mencuci sebagian piring kotor',1000,20000,'test','2023-08-24','2023-07-31 14:08:56','2023-07-31 14:08:56'),
(10,1,'PENCUCI KAKI','mencuci sebagian piring kotor',1000,20000,'test','2023-08-24','2023-07-31 14:09:02','2023-07-31 14:09:02'),
(11,1,'PENCUCI KAKI','mencuci sebagian piring kotor',1000,20000,'test','2023-08-24','2023-07-31 14:22:09','2023-07-31 14:22:09');

/*Table structure for table `pendidikan_formal_users` */

DROP TABLE IF EXISTS `pendidikan_formal_users`;

CREATE TABLE `pendidikan_formal_users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_user` int(11) NOT NULL,
  `nama_sekolah` varchar(255) DEFAULT NULL,
  `tanggal_masuk` date DEFAULT NULL,
  `tanggal_lulus` date DEFAULT NULL,
  `tingkat_pendidikan` varchar(255) DEFAULT NULL,
  `jurusan` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

/*Data for the table `pendidikan_formal_users` */

insert  into `pendidikan_formal_users`(`id`,`id_user`,`nama_sekolah`,`tanggal_masuk`,`tanggal_lulus`,`tingkat_pendidikan`,`jurusan`,`created_at`,`updated_at`) values 
(1,4,'SMA SUKA SUKA','2022-07-31','2023-07-31','SMA','IPA','2023-07-31 09:34:32','2023-07-31 09:40:35'),
(2,4,'UNP (UNIVERSITAS PERSIB)','2022-07-31','2023-07-31','S1','IT','2023-07-31 09:35:05','2023-07-31 09:40:56');

/*Table structure for table `pendidikan_non_formal_users` */

DROP TABLE IF EXISTS `pendidikan_non_formal_users`;

CREATE TABLE `pendidikan_non_formal_users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_user` int(11) NOT NULL,
  `nama_sekolah` varchar(255) DEFAULT NULL,
  `tanggal_masuk` date DEFAULT NULL,
  `tanggal_lulus` date DEFAULT NULL,
  `jurusan` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

/*Data for the table `pendidikan_non_formal_users` */

insert  into `pendidikan_non_formal_users`(`id`,`id_user`,`nama_sekolah`,`tanggal_masuk`,`tanggal_lulus`,`jurusan`,`created_at`,`updated_at`) values 
(1,4,'BOTCAMP UCUP NIRIN','2022-07-31','2023-07-31','DEV OPS','2023-07-31 09:45:09','2023-07-31 09:45:09'),
(2,4,'BOTCAMP WAKWAW','2022-07-31','2023-07-31','FULLSTACK MERN','2023-07-31 09:45:20','2023-07-31 09:45:20');

/*Table structure for table `pengalaman_users` */

DROP TABLE IF EXISTS `pengalaman_users`;

CREATE TABLE `pengalaman_users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_user` int(11) DEFAULT NULL,
  `nama_perusahaan` varchar(255) DEFAULT NULL,
  `tanggal_masuk` date DEFAULT NULL,
  `tanggal_keluar` date DEFAULT NULL,
  `posisi_terakhir` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

/*Data for the table `pengalaman_users` */

insert  into `pengalaman_users`(`id`,`id_user`,`nama_perusahaan`,`tanggal_masuk`,`tanggal_keluar`,`posisi_terakhir`,`created_at`,`updated_at`) values 
(1,4,'PT WAKWAW','2022-07-31','2023-07-31','MOBILE DEVELOPER','2023-07-31 09:47:24','2023-07-31 09:47:24'),
(2,4,'PT UCUP','2022-07-31','2023-07-31','WEB DEVELOPER','2023-07-31 09:47:36','2023-07-31 09:47:36');

/*Table structure for table `perusahaans` */

DROP TABLE IF EXISTS `perusahaans`;

CREATE TABLE `perusahaans` (
  `id_company` int(11) NOT NULL AUTO_INCREMENT,
  `nama` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `role` varchar(255) DEFAULT 'company',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id_company`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

/*Data for the table `perusahaans` */

insert  into `perusahaans`(`id_company`,`nama`,`email`,`password`,`role`,`created_at`,`updated_at`) values 
(1,'PT.TUNA','tuna@mail.com','$argon2id$v=19$m=65536,t=3,p=2$ABwxT2s4fZza4dChLc2Txg$36PPxX4iM0EjpAH1L2Ofa3Aftbkr3vEsVSL/S29UTZk','company','2023-07-31 11:01:25','2023-07-31 11:01:25');

/*Table structure for table `requirement_lowongan_perusahaans` */

DROP TABLE IF EXISTS `requirement_lowongan_perusahaans`;

CREATE TABLE `requirement_lowongan_perusahaans` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_lowongan` int(11) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `deskripsi` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4;

/*Data for the table `requirement_lowongan_perusahaans` */

insert  into `requirement_lowongan_perusahaans`(`id`,`id_lowongan`,`nama`,`deskripsi`) values 
(3,7,'semua gender','semua gender bisa, asal tekun'),
(4,7,'wfh','bekerja di rumah masing-masing'),
(7,9,'semua gender','semua gender bisa, asal tekun'),
(8,9,'wfh','bekerja di rumah masing-masing'),
(9,10,'semua gender','semua gender bisa, asal tekun'),
(10,10,'wfh','bekerja di rumah masing-masing'),
(11,11,'semua gender','semua gender bisa, asal tekun'),
(12,11,'wfh','bekerja di rumah masing-masing');

/*Table structure for table `users` */

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id_user` int(11) NOT NULL AUTO_INCREMENT,
  `first_name` varchar(255) DEFAULT NULL,
  `last_name` varchar(255) DEFAULT NULL,
  `username` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `specialist` varchar(255) DEFAULT NULL,
  `role` varchar(255) DEFAULT 'user',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id_user`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

/*Data for the table `users` */

insert  into `users`(`id_user`,`first_name`,`last_name`,`username`,`email`,`password`,`specialist`,`role`,`created_at`,`updated_at`) values 
(2,'heii','test','test','test1@mail.com','$argon2id$v=19$m=65536,t=3,p=2$14axJPPsJfJC0gPIAsDTgw$0qnzngAqPsKJcRdgJsIMr+HAW/pL4P987Uek7U136cA','test','user','2023-07-29 15:22:54','2023-07-29 15:22:54'),
(3,'heii','test','test','test12@mail.com','$argon2id$v=19$m=65536,t=3,p=2$KNZuLE4JUIHEYQt3w4wfQg$GGhyLNab6hGU/miI+psp4R0XKPe9dLjJmuk7GfbKQJc','test','user','2023-07-29 15:23:05','2023-07-29 15:23:05'),
(4,'herdiyana','firmansyah','herdin','herdiyana@mail.com','$argon2id$v=19$m=65536,t=3,p=2$TrwlAwCexgDruLEqiWwTqw$sRiJiRC828poBAcGakZASpKPhXbCYw4+9rg3g5Qs41s','tester','user','2023-07-31 09:19:29','2023-07-31 09:19:29');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
