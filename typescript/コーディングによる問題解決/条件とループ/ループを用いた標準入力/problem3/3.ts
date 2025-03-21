// 問題文
// N 個の正の整数 A0,A1,…,AN-1が与えられます。N 個の整数のうち、3 の倍数であるものを改行区切りで全て出力してください。
// また、入力される値は次の制約を満たします。
// N は 1 以上 1000 以下の整数
// Ai​は 1 以上 1000 以下の整数 (0≤i≤N−1)

const MIN_VALUE = 1;
const MAX_VALUE = 1000;

/** Nの数とAの要素数が一致しているか確認 */
const checkLength = (N: number, A: number[]): boolean => {
    return N === A.length
}

/** 1 以上 1000 以下の整数であるか確認 */
const checkNumberInRange = (num: number): boolean => {
    return num >= MIN_VALUE && num <= MAX_VALUE
}

/** 配列の要素が全て 1 以上 1000 以下の整数であるか確認 */
const checkArrayNumbersInRange = (arr: number[]): boolean => {
    return arr.every(num => checkNumberInRange(num))
}

/**
 * N 個の正の整数 A0, A1, …, AN-1 のうち、3 の倍数であるものを改行区切りで全て出力する
 * @param {number} N - 配列の要素数
 * @param {number[]} A - 数値の配列
 */
const outputMultiplesOfThree = (N: number, A: number[]): void => {
    if (!checkNumberInRange(N)) {
        console.log(`Nは${MIN_VALUE}以上${MAX_VALUE}以下の整数を入力してください`);
        return
    }

    if (!checkLength(N, A)) {
        console.log("Nの値と配列の要素数が一致しません");
        return
    }

    if (!checkArrayNumbersInRange(A)) {
        console.log(`配列の要素は${MIN_VALUE}以上${MAX_VALUE}以下の整数を入力してください`);
        return
    }

    for(let i: number = 0; i < A.length; i++) {
        if(A[i] % 3 === 0) {
            console.log(A[i])
        }
    }
}

outputMultiplesOfThree(3, [27, 18, 28]);
outputMultiplesOfThree(3, [31, 41, 59]);
outputMultiplesOfThree(2, [10, 20, 30]); // エラーケース
outputMultiplesOfThree(1001, [10, 20]); // エラーケース
outputMultiplesOfThree(3, [0, 20, 30]); // エラーケース