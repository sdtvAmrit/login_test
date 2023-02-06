from airflow import DAG
from airflow.operators.python_operator import PythonOperator
from datetime import datetime, timedelta

# Fetch the IP addresses of Adder and Printer from Consul
def fetch_ip_from_consul(**kwargs):
    # Fetch the IP addresses of Adder and Printer from Consul here

    return (adder_ip, printer_ip)

# Call Adder microservice and pass the sum to Printer microservice
def call_adder_and_printer(**kwargs):
    # Call Adder microservice to get the sum
    # Pass the sum to Printer microservice

    return

default_args = {
    'owner': 'airflow',
    'depends_on_past': False,
    'start_date': datetime(2023, 2, 5),
    'retries': 1,
    'retry_delay': timedelta(minutes=5)
}

dag = DAG(
    'adder_printer_dag',
    default_args=default_args,
    description='A DAG to send sum from Adder to Printer',
    schedule_interval=timedelta(hours=1)
)

fetch_ip_task = PythonOperator(
    task_id='fetch_ip_from_consul',
    python_callable=fetch_ip_from_consul,
    dag=dag
)

call_adder_printer_task = PythonOperator(
    task_id='call_adder_and_printer',
    python_callable=call_adder_and_printer,
    dag=dag,
    provide_context=True
)

fetch_ip_task >> call_adder_printer_task
