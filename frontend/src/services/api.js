import axios from 'axios'

// Set base url to axios instance
const axiosInstance = () => axios.create({
    baseURL: process.env.VUE_APP_BASE_URL,
});


/*
Here we have all requests to backend
There are urls by request and request functions separated by API publicity
Private API functions have axios config with headers to send JWT token
*/

const urls = {
    login: 'api/public/login/',
    cards: 'api/cards/',
    userInfo: 'api/user/',
    userChangeLang: 'api/user/lang/',
    cardImageUrl: 'api/public/bot-image/',
    processAnswer: 'api/user/answer/',
    leaderboard: `api/leaderboard/`
}

export function publicApi() {
    return {
        login(body) {
            return axiosInstance().post(urls.login, body);
        }
    }
}

export function privateApi() {
    return {
        getUser(jwt) {
            let config = {
                headers: {
                    jwt: jwt
                }
            }
            return axiosInstance().get(urls.userInfo, config);
        },
        changeUserLang(jwt, lang) {
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
        },
        getCards(jwt) {
            let config = {
                headers: {
                    jwt: jwt
                }
            }
            return axiosInstance().get(urls.cards, config);
        }
    }
}