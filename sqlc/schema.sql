CREATE TABLE `data_print` (
  `ID` bigint(20) NOT NULL AUTO_INCREMENT,
  `Data_Print_ID` int(11) DEFAULT NULL,
  `Parent_ID` int(11) DEFAULT NULL,
  `Serial_Level` int(11) DEFAULT NULL,
  `Barcode` varchar(200) DEFAULT NULL,
  `Serialisasi` varchar(100) NOT NULL,
  `Batch_No` varchar(100) DEFAULT NULL,
  `Process_Order` varchar(100) NOT NULL,
  `Scanned` datetime DEFAULT NULL,
  `product` varchar(200) DEFAULT NULL,
  `nie` varchar(255) DEFAULT NULL,
  `sku` varchar(255) DEFAULT NULL,
  `Counter` int(11) DEFAULT NULL,
  `Berat` decimal(10,3) DEFAULT NULL,
  `md` datetime DEFAULT NULL,
  `ed` datetime DEFAULT NULL,
  `Username` varchar(200) DEFAULT NULL,
  `station_name` varchar(200) DEFAULT NULL,
  `grup` varchar(200) DEFAULT NULL,
  `ipc` varchar(200) DEFAULT NULL,
  `Sample` datetime DEFAULT NULL,
  `packer` varchar(200) DEFAULT NULL,
  `upload_line` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `status` enum('produksi','karantina','reject','lulus') NOT NULL DEFAULT 'produksi',
  `sjp` varchar(100) DEFAULT NULL,
  `done` tinyint(4) DEFAULT '0',
  `synced` tinyint(4) DEFAULT '0',
  PRIMARY KEY (`ID`) USING BTREE,
  UNIQUE KEY `KEY` (`Data_Print_ID`,`station_name`) USING BTREE,
  KEY `BARCODES` (`Barcode`) USING BTREE,
  KEY `SERIALISASIS` (`Serialisasi`) USING BTREE,
  KEY `LEVELS` (`Serial_Level`) USING BTREE,
  KEY `NIES` (`nie`) USING BTREE,
  KEY `SKUS` (`sku`) USING BTREE,
  KEY `PARENTS` (`Parent_ID`) USING BTREE,
  KEY `SJPS` (`sjp`) USING BTREE,
  KEY `USERNAMES` (`Username`) USING BTREE
)