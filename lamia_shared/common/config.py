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
    PoolMaxConnections: str
    PoolMinConnections: str
    PoolHealthCheckPeriod: str


class Server:
    MaxWorkers: int

    def __init__(self):
        self.MaxWorkers = int(env['MAX_WORKERS'])
