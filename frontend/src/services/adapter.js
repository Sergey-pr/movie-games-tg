import { axiosInstance } from './api';

const urls = {
    login: 'api/public/login/',
    cards: 'api/cards/',
}

export function createDynamicString(string, replaceParams) {
    let result = string;
    if (typeof replaceParams === 'object') {
        for (const dynamicKey in replaceParams) {
            result = result.replaceAll(dynamicKey, replaceParams[dynamicKey]);
        }
    } else {
        replaceParams.forEach((dynamicValue, index) => result = result.replaceAll(`[${index}]`, dynamicValue));
    }
    return result;
}

export function useAuth() {
    return {
        login(body) {
            return axiosInstance().post(urls.login, body);
        }
    }
}

export function useCards() {
    return {
        cardsList() {
            return axiosInstance().get(urls.cards);
        }
    }
}
