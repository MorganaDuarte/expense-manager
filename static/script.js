async function send() {
  try {
    const valueToSend = {
      "value_received": parseFloat(document.getElementById('valueReceived').value),
      "date_received": document.getElementById('dateReceived').value,
      "description_received": document.getElementById('descriptionReceived').value,
      "account_received": document.getElementById('accountReceived').value,
    };

    const response = await fetch('/api/values-received', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(valueToSend),
    });

    if (!response.ok) throw new Error('Error calling API');

    const data = await response.json();
    document.getElementById('value').innerText = data.value;
    document.getElementById('date').innerText = data.date;
    document.getElementById('description').innerText = data.description;
    document.getElementById('account').innerText = data.account;
  } catch (error) {
    console.error(error);
    document.getElementById('value').innerText = 'Erro na requisição';
  }
}