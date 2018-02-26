<?php
$username = "dsadmin"
$password = "dsadmin789"
$hostname = "dsa-instance1.cbjkxisxciuj.us-east-2.rds.amazonaws.com"
$dbname = "dsalab5db"

//Conexão para o banco de dados
$dbhandle = mysql_connect($hostname, $username, $password) or die("Conection failed");
echo "Conectado ao MySQL! Username - $username, Senha = $password, host - $hostname<br>";
$selected = mysql_select_db("$dbname",$dbhandle) or die("Falha ao conectar o DB");
?>