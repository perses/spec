// Copyright The Perses Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { DurationString } from '@perses-dev/spec';

export interface SQLProxy {
  kind: 'sqlproxy';
  spec: SQLProxySpec;
}

export interface SQLProxySpec {
  driver: 'mysql' | 'mariadb' | 'postgres';
  // host is the hostname required to contact the datasource
  host: string;
  // database is the database for the datasource
  database: string;
  // secret is the name of the secret that should be used for the proxy or discovery configuration
  // It will contain any sensitive information such as username, password, token, certificate.
  secret?: string;
  // MySQL specific driver config
  mysql?: MySQLConfig;
  // MariaDB specific driver config (uses same structure as MySQL since MariaDB is MySQL-compatible)
  mariadb?: MySQLConfig;
  // Postgres specific driver config
  postgres?: PostgresConfig;
}

export interface MySQLConfig {
  params?: Record<string, string>;
  maxAllowedPacket?: number;
  timeout?: DurationString;
  readTimeout?: DurationString;
  writeTimeout?: DurationString;
}

export interface PostgresConfig {
  // maxConns is the maximum size of the pool
  maxConns?: number;
  // connectTimeout the timeout value used for socket connect operations.
  connectTimeout?: DurationString;
  // PrepareThreshold specifies the number of PreparedStatement executions that must occur before the driver begins using server-side prepared statements.
  prepareThreshold?: DurationString;
  sslMode?: 'disable' | 'allow' | 'prefer' | 'require' | 'verify-full' | 'verify-ca';
}
