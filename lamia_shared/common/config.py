from os import environ as env


class Config:

    def __init__(self):
        self.HTTP = HTTP()
        self.Server = Server()
        self.Database = Database()


class HTTP:
    Host: str
    Port: str

    def __init__(self):
        self.Host = env['HOST']
        self.Port = env['PORT']


class Database:
    Name: str
    Host: str
    Port: str
    Username: str
    Password: str
    PoolMaxConnections: int
    PoolMinConnections: int
    PoolHealthCheckPeriod: str

    def __init__(self):
        self.Name = env['PG_DB']
        self.Host = env['PG_HOST']
        self.Port = env['PG_PORT']
        self.Username = env['PG_USER']
        self.Password = env['PG_PASSWORD']
        self.PoolMaxConnections = int(env['PG_POOL_MAX_CONNECTIONS'])
        self.PoolMinConnections = int(env['PG_POOL_MIN_CONNECTIONS'])
        self.PoolHealthCheckPeriod = env['PG_POOL_HEALTH_CHECK_PERIOD']


class Server:
    MaxWorkers: int

    def __init__(self):
        self.MaxWorkers = int(env['MAX_WORKERS'])
