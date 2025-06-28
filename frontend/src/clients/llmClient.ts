import type { OllamaBody, OllamaResponse } from "@/types";
import ky from 'ky';

const baseUrl = new URL("http://localhost:8080/model");

export const questionOllama = async (body: OllamaBody): Promise<any> => {
    try {
        const resp = await ky.post(baseUrl, {
            json: body,
            headers: {
                "Content-Type": "application/json",
            }
        }).json();
        
        console.log({resp});
        return resp;
    } catch (error) {
        console.error('Error calling Ollama:', error);
        throw error;
    }
};
