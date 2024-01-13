

export const generateRandomString = () => {
    const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    const length = 10; 
    let result = '';
  
    for (let i = 0; i < length; i++) {
    const randomIndex = Math.floor(Math.random() * characters.length);
    result += characters.charAt(randomIndex);
    }
  
    return result
};


export function generateArrayOfNumbers() {
    const arrayOfNumbers = [];

    for (let i = 0.5; i < 1.6; i += 0.1) {
        const roundedNumber = parseFloat(i.toFixed(1));
        arrayOfNumbers.push(roundedNumber);
    }

    return arrayOfNumbers;
}