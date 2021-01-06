#!/bin/bash
mysqldump -uwebgo_user -pwebgo_password --databases webgo_db > /backup/backup.sql 