--- 

# Backend Application

The Backend Application is a part of the ZmeyGorynych Project. It integrates with both SQL and NoSQL databases and uses Redis for caching and key-value storage.

## Components

### Backend Application (`zg_backend`)
This component serves as the backend for the project, handling data operations and integrating with various databases.

#### Docker Compose Configuration
```yaml
version: '3'

networks:
  local-net:
    external: true

services:

  zg_backend:
    build:
      context: .
      dockerfile: ./Dockerfile
    container_name: zg_backend
    env_file:
      - .env-docker
    networks:
      - local-net
    volumes:
      - .:/app
    ports:
      - "8080:8080"
```

#### Environment Variables (`.env-docker`)
```env
MYSQL_DB_1_HOST=zg_mysql_db_1
MYSQL_DB_2_HOST=zg_mysql_db_2
MYSQL_DB_1_PORT=3306
MYSQL_DB_2_PORT=3307
MYSQL_ROOT_PASSWORD=rootpassword
MYSQL_DATABASE=db
MYSQL_USER=user
MYSQL_PASSWORD=password
ZG_NOSQL_REDIS_URL=zg_nosql_redis_index:6379
ZG_SQL_REDIS_URL=zg_sql_repo_redis:16379
ZG_REDIS_DB=0
ZG_REDIS_CACHE_DB=1
ZG_REDIS_EXP_TIME=3600s
ZG_BACKEND_PORT=8080
ZG_MONGODB_URL_1=mongodb://zg_mongodb:27017/zg_mongodb
ZG_MONGODB_URL_2=mongodb://zg_mongodb_2:27017/zg_mongodb_2
```

#### Configuration File (`config.yaml`)
```yaml
telemetry:
  enabled: true
  interval: 10
  address: localhost:8888

sql_dbs:
  - host: ${MYSQL_DB_1_HOST}
    port: ${MYSQL_DB_1_PORT}
    database: ${MYSQL_DATABASE}
    user: ${MYSQL_USER}
    password: ${MYSQL_PASSWORD}

  - host: ${MYSQL_DB_2_HOST}
    port: ${MYSQL_DB_2_PORT}
    database: ${MYSQL_DATABASE}
    user: ${MYSQL_USER}
    password: ${MYSQL_PASSWORD}

nosql_dbs:
  mongodb:
    - ${ZG_MONGODB_URL_1}
    - ${ZG_MONGODB_URL_2}

sql_cache:
  address: ${ZG_SQL_REDIS_URL}
  db: ${ZG_REDIS_CACHE_DB}
  exp_time: ${ZG_REDIS_EXP_TIME}

sql_kv_db:
  address: ${ZG_SQL_REDIS_URL}
  db: ${ZG_REDIS_DB}

nosql_kv_db:
  address: ${ZG_NOSQL_REDIS_URL}
  db: ${ZG_REDIS_DB}

server:
  port: ${ZG_BACKEND_PORT}
```

## Getting Started

### Prerequisites
- Docker
- Docker Compose

### Running the Backend Application
1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/message-generator.git
   cd message-generator/backend
   ```
2. Build and run the Docker containers:
   ```bash
   docker-compose up --build
   ```

### Environment Variables
Ensure to set the following environment variables in the `.env-docker` file:
- `MYSQL_DB_1_HOST`: Hostname of the first MySQL instance (e.g., `zg_mysql_db_1`).
- `MYSQL_DB_2_HOST`: Hostname of the second MySQL instance (e.g., `zg_mysql_db_2`).
- `MYSQL_DB_1_PORT`: Port of the first MySQL instance (e.g., `3306`).
- `MYSQL_DB_2_PORT`: Port of the second MySQL instance (e.g., `3307`).
- `MYSQL_ROOT_PASSWORD`: Root password for MySQL instances (e.g., `rootpassword`).
- `MYSQL_DATABASE`: Database name (e.g., `db`).
- `MYSQL_USER`: Username for MySQL instances (e.g., `user`).
- `MYSQL_PASSWORD`: Password for MySQL instances (e.g., `password`).
- `ZG_NOSQL_REDIS_URL`: URL of the NoSQL Redis server (e.g., `zg_nosql_redis_index:6379`).
- `ZG_SQL_REDIS_URL`: URL of the SQL Redis server (e.g., `zg_sql_repo_redis:16379`).
- `ZG_REDIS_DB`: Redis database number for KV store (e.g., `0`).
- `ZG_REDIS_CACHE_DB`: Redis database number for caching (e.g., `1`).
- `ZG_REDIS_EXP_TIME`: Expiration time for Redis cache (e.g., `3600s`).
- `ZG_BACKEND_PORT`: Port for the backend server (e.g., `8080`).
- `ZG_MONGODB_URL_1`: URL of the first MongoDB instance (e.g., `mongodb://zg_mongodb:27017/zg_mongodb`).
- `ZG_MONGODB_URL_2`: URL of the second MongoDB instance (e.g., `mongodb://zg_mongodb_2:27017/zg_mongodb_2`).

## License
This project is licensed under the MIT License.

---
