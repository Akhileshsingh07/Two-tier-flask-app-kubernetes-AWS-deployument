version: "3"
services: 
    backend:
	  img:
	  ports:
	    - "5000:5000"
      enviroment:
	   MY_HOST: "mysql"
	   MYSQL_USER: "admin"
	   MYSQL_PASSWORD: "admin"
	   MY_DB: "myDb"
      depends_on:
	   -mysql

	mysql:
	  img: mysql:5.7
	  enviroment:
	   MYSQL_ROOT_PASSWORD: root
       MYSQL_DATABASE: myDb
       MYSQL_USER: admin
       MYSQL_PASSWORD: admin
      ports: 
	   -"3306:3306"

	#forcreatingtables
    volumes:
      - ./message.sql:/docker-entrypoint-initdb.d/message.sql
	  - mysql-data:/var/lib/mysql #bind the hard disk of laptop and containers to persist data

volumes:
   mysql-data: