from setuptools import setup, find_packages

setup(
    name='lamia_shared',
    version='1.0.26',
    description='Lamia Shared Library',
    author='Erfan Sahebi',
    packages=find_packages(),
    install_requires=[
        'protobuf~=4.23.3',
        'grpcio~=1.56.0',
        'grpcio-tools~=1.56.0'
    ]
)
