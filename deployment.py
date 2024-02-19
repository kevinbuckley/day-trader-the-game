import subprocess
import json
import requests
from dotenv import load_dotenv
import os


def run_validate() -> bool:
    print('STARTING: terraform validate')
    result = subprocess.run(['terraform', 'validate'], cwd='./terraform', capture_output=True, text=True)
    if result.returncode == 0:
        print('Terraform is valid')
        return True
    else:
        print(result.stderr)
        return False
    return True

def run_apply():
    print('STARTING: terraform apply')
    result = subprocess.run(['terraform', 'apply', '-auto-approve'], cwd='./terraform', capture_output=True, text=True)
    if(result.returncode == 0):
        print(f'apply successful {result.stdout}')
    else:
        print(f'apply failure: {result.stderr}')


def call_api():
    requests.post()
    url = "https://api.example.com"
    headers = {
        "Content-Type": "application/json",
        "Authorization": "Bearer <your_token>"
    }
    data = {
        "param1": "value1",
        "param2": "value2"
    }
    response = requests.post(url, headers=headers, json=data)
    if response.status_code == 200:
        print("API call successful")
    else:
        print(f"API call failed with status code {response.status_code}")

def read_resources():
    file_name = './terraform/terraform.tfstate'

    if(os.access(file_name, os.R_OK) is False):
        print(f'this script can not access {file_name}')

    if os.path.exists(file_name):
        state = json.loads(file_name)
        for resources in state['resources']:
            print(f'[{resources["type"]} :: {resources["name"]}]\n')
    else:
        print('No state is managed at this time')


if __name__ == '__main__':
    load_dotenv(dotenv_path='./terraform/.env')
    read_resources()
    valid = run_validate()
    if(valid):
        run_apply()