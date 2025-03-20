// 問題文
// N 個の正の整数 A0,A1,…,AN-1が与えられます。N 個の整数の合計値を求めてください。
// また、入力される値は次の制約を満たします。
// N は 1 以上 1000 以下の整数
// Ai​は 1 以上 1000 以下の整数 (0≤i≤N−1)

const MIN_VALUE = 1;
const MAX_VALUE = 1000;

/** Nの数とAの要素数が一致しているか確認 */
const checkLength = (N: number, A: number[]): boolean => {
    return N === A.length;
};

/** 1 以上 1000 以下の整数であるか確認 */
const checkNumberInRange = (num: number): boolean => {
    return num >= MIN_VALUE && num <= MAX_VALUE;
};

/** 配列の要素が全て 1 以上 1000 以下の整数であるか確認 */
const checkArrayNumbersInRange = (arr: number[]): boolean => {
    return arr.every(num => checkNumberInRange(num));
};

/** 配列の合計を求める */
const sumOfArray = (arr: number[]): number => {
    return arr.reduce((sum, element) => sum + element, 0);
};

/** 答えを計算して表示 */
const calculateAndPrintResult = (N: number, A: number[]): number | string => {
    if (!checkNumberInRange(N)) {
        return `Nは${MIN_VALUE}以上${MAX_VALUE}以下の整数を入力してください`;
    }

    if (!checkLength(N, A)) {
        return "Nの値と配列の要素数が一致しません";
    }

    if (!checkArrayNumbersInRange(A)) {
        return `配列の要素は${MIN_VALUE}以上${MAX_VALUE}以下の整数を入力してください`;
    }

    return sumOfArray(A);
};

console.log(calculateAndPrintResult(3, [10, 20, 30]));
console.log(calculateAndPrintResult(2, [10, 20, 30])); // エラーケース
console.log(calculateAndPrintResult(1001, [10, 20])); // エラーケース
console.log(calculateAndPrintResult(3, [0, 20, 30])); // エラーケース