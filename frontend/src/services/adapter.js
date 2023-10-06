import { axiosInstance } from './api';

const urls = {
    login: 'api/public/login/',
    cards: 'api/cards/',
    userInfo: 'api/user/',
    userChangeLang: 'api/user/lang/',
    cardImageUrl: 'api/public/bot-image/',
    processAnswer: 'api/user/answer/',
    leaderboard: `api/leaderboard/`
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
        },
        processAnswer(jwt, points, cardId) {
            let config = {
                headers: {
                    jwt: jwt
                }
            }
            return axiosInstance().post(urls.processAnswer, {"points": points, card_id: cardId}, config);
        },
        getLeaderboards(jwt) {
            let config = {
                headers: {
                    jwt: jwt
                }
            }
            return axiosInstance().get(urls.leaderboard, config);
        }
    }
}