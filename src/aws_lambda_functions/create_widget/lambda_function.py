import databases
import utils
import logging
import sys
import os
import uuid
from psycopg2.extras import RealDictCursor


"""
Define the connection to the database outside of the "lambda_handler" function.
The connection to the database will be created the first time the function is called.
Any subsequent function call will use the same database connection.
"""
postgresql_connection = None

# Define databases settings parameters.
POSTGRESQL_USERNAME = os.environ["POSTGRESQL_USERNAME"]
POSTGRESQL_PASSWORD = os.environ["POSTGRESQL_PASSWORD"]
POSTGRESQL_HOST = os.environ["POSTGRESQL_HOST"]
POSTGRESQL_PORT = int(os.environ["POSTGRESQL_PORT"])
POSTGRESQL_DB_NAME = os.environ["POSTGRESQL_DB_NAME"]

logger = logging.getLogger(__name__)  # Create the logger with the specified name.
logger.setLevel(logging.WARNING)  # Set the logging level of the logger.


def lambda_handler(event, context):
    """
    :argument event: The AWS Lambda uses this parameter to pass in event data to the handler.
    :argument context: The AWS Lambda uses this parameter to provide runtime information to your handler.
    """
    # Since the connection with the database were defined outside of the function, we create global variable.
    global postgresql_connection
    if not postgresql_connection:
        try:
            postgresql_connection = databases.create_postgresql_connection(
                POSTGRESQL_USERNAME,
                POSTGRESQL_PASSWORD,
                POSTGRESQL_HOST,
                POSTGRESQL_PORT,
                POSTGRESQL_DB_NAME
            )
        except Exception as error:
            logger.error(error)
            sys.exit(1)

    # Define the values of the data passed to the function.
    organization_id = event["arguments"]["input"]["organizationId"]
    channel_name = event["arguments"]["input"]["channelName"]
    channel_description = event["arguments"]["input"].get("channelDescription", None)

    # Generating a short technical identifier for the widget.
    channel_technical_id = utils.ShortUUID().encode(uuid.uuid4())

    # With a dictionary cursor, the data is sent in a form of Python dictionaries.
    cursor = postgresql_connection.cursor(cursor_factory=RealDictCursor)

    # Prepare the SQL request that creates the new channel.
    statement = """
    insert into channels (
        channel_name,
        channel_description,
        channel_type_id,
        channel_technical_id
    )
    values (
        {0},
        {1},
        'ad0b1434-a03c-4de2-8a5c-c22366f2710c',
        '{2}'
    )
    returning channel_id;
    """.format(
        'null' if channel_name is None or len(channel_name) == 0
        else "'{0}'".format(channel_name.replace("'", "''")),
        'null' if channel_description is None or len(channel_description) == 0
        else "'{0}'".format(channel_description),
        channel_technical_id
    )

    # Execute a previously prepared SQL query.
    try:
        cursor.execute(statement)
    except Exception as error:
        logger.error(error)
        sys.exit(1)

    # After the successful execution of the query commit your changes to the database.
    postgresql_connection.commit()

    # Fetch the next row of a query result set.
    channel_id = cursor.fetchone()["channel_id"]

    # Prepare the SQL request that creates the new channel.
    statement = """
    insert into channels_organizations_relationship (
        channel_id,
        organization_id
    )
    values (
        '{0}',
        '{1}'
    );
    """.format(
        channel_id,
        organization_id
    )

    # Execute a previously prepared SQL query.
    try:
        cursor.execute(statement)
    except Exception as error:
        logger.error(error)
        sys.exit(1)

    # After the successful execution of the query commit your changes to the database.
    postgresql_connection.commit()

    # The cursor will be unusable from this point forward.
    cursor.close()

    # Create the response format.
    response = {
        "widgetId": channel_technical_id
    }
    return response
