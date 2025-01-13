import RequestResponse from "./RequestResponse.js";

export default class RequestSender{
  async send(url, method, errorMessage, body) {
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
}