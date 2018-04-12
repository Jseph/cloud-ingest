"""
This script is run while configuring the angular application. The logic on this
file runs before the application is started with npm start.
"""

import os
import string
import sys
import collections


def exit_api_not_set_error():
    if not INGEST_API_URL:
        print('ERROR: OPI_API_URL variable not set. Please set the OPI_API_URL '
              'environment variable to the API url and run again.')
        sys.exit(1)


# TODO: Define different API URLS for perf, prod and test environments.
INGEST_API_URL = ''
try:
    INGEST_API_URL = os.environ['OPI_API_URL']
except:
    exit_api_not_set_error()

if not INGEST_API_URL:
    exit_api_not_set_error()

Environment = collections.namedtuple('Environment',
                                     'filename client_id is_prod')

# The environment files to write.
ENVIRONMENTS = [
    Environment(
        filename='environment.perf.ts',
        client_id=
        '626613183123-o1l2r81kov2fuii2pc1p5e67k6facktt.apps.googleusercontent.com',
        is_prod='true'),
    Environment(
        filename='environment.prod.ts',
        client_id=
        '23880221060-9pa2oef8ko1q8o45mvjfq5d6dfgf1qf0.apps.googleusercontent.com',
        is_prod='true'),
    Environment(
        filename='environment.test.ts',
        client_id=
        '701178595865-por9ijjvgbjoka841c1mkki23tqka66a.apps.googleusercontent.com',
        is_prod='true'),
    Environment(
        filename='environment.ts',
        client_id=
        '416127938080-vosnnsq7758ub1iai84ei3u1enstq8kp.apps.googleusercontent.com',
        is_prod='false')
]

# The directory with the environment files.
ENV_DIRECTORY = 'src/environments/'

# The production values for each one of the environments.
IS_PROD_VALUES = ['true', 'true', 'true', 'false']

TEMPLATE = string.Template("""
export const environment = {
  production: '$IS_PRODUCTION',
  apiUrl: '$API_URL',
  authClientId: '$AUTH_CLIENT_ID',
};
""")

if not os.path.exists(ENV_DIRECTORY):
    os.makedirs(ENV_DIRECTORY)

for environment in ENVIRONMENTS:
    with open(ENV_DIRECTORY + environment.filename, "w") as env_file:
        env_file.write(
            TEMPLATE.substitute(
                IS_PRODUCTION=environment.is_prod,
                API_URL=INGEST_API_URL,
                AUTH_CLIENT_ID=environment.client_id))
