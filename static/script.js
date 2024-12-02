async function clickOnMe() {
  try {
    const valueToSend = {
      "value_received": document.getElementById('valueReceived').value,
      "date_received": document.getElementById('dateReceived').value
    };

    const response = await fetch('/api/values-received', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(valueToSend),
    });

    if (!response.ok) throw new Error('Error calling API');

    const data = await response.json();
    console.log(data);
    document.getElementById('value-received').innerText = data.value
    document.getElementById('date-received').innerText = data.date
  } catch (error) {
    console.error(error);
    document.getElementById('value-received').innerText = 'Erro na requisição';
  }
}
