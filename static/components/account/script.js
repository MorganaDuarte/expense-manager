async function saveAccount(event) {
  event.preventDefault();

  try {
    const valueToSend = {
      bank_value: document.getElementById('bankValue').value.trim(),
      account_value: document.getElementById('accountValue').value.trim(),
      acronym_value: document.getElementById('acronymValue').value.trim(),
    };

    if (!valueToSend.bank_value || !valueToSend.account_value || !valueToSend.acronym_value) {
      throw new Error('Todos os campos são obrigatórios!');
    }

    const response = await fetch('/api/save-account', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(valueToSend),
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Erro ao salvar os dados. Por favor, tente novamente.');
    }


    document.getElementById('bankSave').innerText = valueToSend.bank_value;
    document.getElementById('accountSave').innerText = valueToSend.account_value;
    document.getElementById('acronymSave').innerText = valueToSend.acronym_value;

    document.getElementById('accountForm').reset();

    document.getElementById('errorMessage').innerText = '';
  } catch (error) {
    console.error(error);
    document.getElementById('errorMessage').innerText = error.message;
  }
}
