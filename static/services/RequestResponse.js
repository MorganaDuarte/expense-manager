export class RequestResponse{
  constructor(body, error) {
    this.body = body
    this.error = error
  }

  async sendRequest(url, method, errorMessage, body) {
    let bodyResponse
    let errorResponse

    try {
      const response = await fetch(url, {
        method: method,
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(body)
      });

      if (!response.ok) throw new Error(errorMessage);

      bodyResponse = await this.processResponse(response);
    } catch (error) {
      console.error(error);
      errorResponse = error
    }

    return new RequestResponse(bodyResponse, errorResponse);
  }

  async processResponse(response) {
    const contentType = response.headers.get('Content-Type');

    return contentType?.includes('application/json') ? await response.json() : await response.text();
  }

  hasError() {
    return !!this.error;
  }

  getErrorMessage() {
    return this.error.message;
  }

  getBody() {
    return this.body;
  }
}