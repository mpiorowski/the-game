import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";

const firebaseConfig = {
    apiKey: "AIzaSyDuNGWqFiT_7rWAm_wWsdHGkBvXEv9bIVA",
    authDomain: "homeit-app-364020.firebaseapp.com",
};

export const app = initializeApp(firebaseConfig);
export const auth = getAuth(app);
