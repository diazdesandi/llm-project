import type { ModelBody, ModelResponse } from "@/types";
import ky from 'ky';

const baseUrl = new URL("http://localhost:8080/model");

export const questionOllama = async (body: ModelBody): Promise<ModelResponse> => {
    try {
        const resp = await ky.post(baseUrl, {
            json: body,
            headers: {
                "Content-Type": "application/json",
            }
        }).json();
        
        return resp as ModelResponse
    } catch (error) {
        console.error('Error calling Ollama:', error);
        throw error;
    }
};
