import { axiosInstance } from './api';

const urls = {
    login: 'api/public/login/',
    cards: 'api/cards/',
    userInfo: 'api/user/',
    userChangeLang: 'api/user/lang/'
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
        cardsList(jwt) {
            let config = {
                headers: {
                    jwt: jwt
                }
            }
            return axiosInstance().get(urls.cards, config);
        }
    }
}

export function useUsers() {
    return {
        user(jwt) {
            let config = {
                headers: {
                    jwt: jwt
                }
            }
            return axiosInstance().get(urls.userInfo, config);
        },
        changeLang(jwt, lang) {
            let config = {
                headers: {
                    jwt: jwt
                }
            }
            return axiosInstance().post(urls.userChangeLang, {"language_code": lang}, config);
        }
    }
}