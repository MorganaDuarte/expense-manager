async function saveAccount() {
  try {
    const valueToSend = {
      "bank_value": document.getElementById('bankValue').value,
      "account_value": document.getElementById('accountValue').value,
      "acronym_value": document.getElementById('acronymValue').value,
    };

    const response = await fetch('/api/save-account', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(valueToSend),
    });

    if (!response.ok) throw new Error('Error calling API');

    document.getElementById('bankSave').innerText = valueToSend.bank_value;
    document.getElementById('accountSave').innerText = valueToSend.account_value;
    document.getElementById('acronymSave').innerText = valueToSend.acronym_value;
  } catch (error) {
    console.error(error);
    document.getElementById('value').innerText = 'Erro na requisição';
  }
}