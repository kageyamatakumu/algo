// 問題文
// N 個の正の整数 S0,S1,…,SN-1が与えられます。N 個の文字列を前からすべてつなげてできる文字列の長さを出力してください。
// また、入力される値は次の制約を満たします。
// N は 1 以上 1000 以下の整数
// 英小文字からなる長さ 1 以上 2000 以下の文字列 (0≤i≤N−1)

const MIN_VALUE = 1;
const MAX_VALUE = 1000;

const MIN_LENGTH = 1;
const MAX_LENGTH = 2000;


/** Nの数とSの要素数が一致しているか確認 */
const checkLength = (N: number, S: string[]): boolean => {
    return N === S.length
}

/** 1 以上 1000 以下の整数であるか確認 */
const checkNumberInRange = (num: number): boolean => {
    return num >= MIN_VALUE && num <= MAX_VALUE
}


/** 1 以上 2000 以下の文字列であるか確認 */
const checkStringLengthInRange = (str: string): boolean => {
    const length: number = str.length;
    return length >= MIN_LENGTH && length <= MAX_LENGTH
}

/** 配列の要素（文字列）が全て 1 以上 2000 以下の長さであるか確認 */
const checkArrayStringLengthInRange = (arrStr: string[]): boolean => {
    return arrStr.every(element => checkStringLengthInRange(element))
}

/**
 * N 個の正の整数 S0,S1,…,SN-1が与えられます。N 個の文字列を前からすべてつなげてできる文字列の長さを出力してください。
 * @param {number} N - 配列の要素数
 * @param {string[]} arrS - 文字列の配列
 */
const calculateTotalStringLength = (N: number, arrS: string[]): void => {

    if (!checkNumberInRange(N)) {
        console.log(`Nは${MIN_VALUE}以上${MAX_VALUE}以下の整数を入力してください`);
        return
    }

    if (!checkLength(N, arrS)) {
        console.log("Nの値と配列の要素数が一致しません");
        return
    }

    if (!checkArrayStringLengthInRange(arrS)) {
        console.log(`文字列は${MIN_LENGTH}以上${MAX_LENGTH}以下で入力してください`);
        return
    }

    // 配列の文字を結合する
    const joinStr: string = arrS.join("")

    console.log(joinStr.length)
}

calculateTotalStringLength(3, ["hello", "algo", "method"]);
calculateTotalStringLength(2, ["hello", "algo", "method"]); //エラーケース
calculateTotalStringLength(1001, ["hello", "algo", "method"]); //エラーケース
calculateTotalStringLength(1, ["a".repeat(20001)])

