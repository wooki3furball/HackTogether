# Author: Michael Bopp
# Filename: script.py
# Description: Python program to configure Postgres Dockerfile with build & run commands.
# Date Created: 3/02/23
# Last Edited: 3/02/23
# Dependency: Python Interpreter

import os
import subprocess
import sys
import argparse

def docker_build(no_cache=False):
    """
    Construct the Docker build command for PostgreSQL
    """
    docker_build_command = [
        'docker', 'build',
        '-t', 'postgres:toegetherTest',  # Ensure this tag matches in both build and run functions
        '.'
    ]
    
    if no_cache:
        docker_build_command.append('--no-cache')
    
    completed_process = subprocess.run(docker_build_command, capture_output=True)
    if completed_process.returncode != 0:
        print("Docker build failed:", completed_process.stderr)
        sys.exit(1)

def docker_run(port_mapping):
    """
    Construct the Docker run command for PostgreSQL
    """
    docker_run_command = [
        'docker', 'run',
        '--name', 'Postgres',
        '-p', f'{port_mapping}:5432',
        '-e', f'POSTGRES_PASSWORD={os.getenv("POSTGRES_PASSWORD", "Test!Pass123")}',  # Use environment variable or default
        '-e', 'POSTGRES_DB=exampledb',
        '-e', 'POSTGRES_USER=exampleuser',
        '-d', 'postgres:toegetherTest'  # Ensure this tag matches the one used in docker_build
    ]
    
    completed_process = subprocess.run(docker_run_command, capture_output=True)
    if completed_process.returncode != 0:
        print("Docker run failed:", completed_process.stderr)
        sys.exit(1)

if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument("action", choices=["build", "run"], help="Either 'build', or 'run'")
    parser.add_argument("--no-cache", action="store_true", help="Use --no-cache with Docker build")
    parser.add_argument("--port", default="5432", help="Port mapping for PostgreSQL container")
    args = parser.parse_args()

    if args.action == 'build':
        docker_build(no_cache=args.no_cache)
    elif args.action == 'run':
        docker_run(args.port)
