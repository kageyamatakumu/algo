// 問題文
// N 個の正の整数 S0,S1,…,SN-1が与えられます。N 個の文字列を前からすべてつなげてできる文字列の長さを出力してください。
// また、入力される値は次の制約を満たします。
// N は 1 以上 1000 以下の整数
// 英小文字からなる長さ 1 以上 2000 以下の文字列 (0≤i≤N−1)
var MIN_VALUE = 1;
var MAX_VALUE = 1000;
var MIN_LENGTH = 1;
var MAX_LENGTH = 2000;
/** Nの数とSの要素数が一致しているか確認 */
var checkLength = function (N, S) {
    return N === S.length;
};
/** 1 以上 1000 以下の整数であるか確認 */
var checkNumberInRange = function (num) {
    return num >= MIN_VALUE && num <= MAX_VALUE;
};
/** 1 以上 2000 以下の文字列であるか確認 */
var checkStringLengthInRange = function (str) {
    var length = str.length;
    return length >= MIN_LENGTH && length <= MAX_LENGTH;
};
/** 配列の要素（文字列）が全て 1 以上 2000 以下の長さであるか確認 */
var checkArrayStringLengthInRange = function (arrStr) {
    return arrStr.every(function (element) { return checkStringLengthInRange(element); });
};
/**
 * N 個の正の整数 S0,S1,…,SN-1が与えられます。N 個の文字列を前からすべてつなげてできる文字列の長さを出力してください。
 * @param {number} N - 配列の要素数
 * @param {string[]} arrS - 文字列の配列
 */
var calculateTotalStringLength = function (N, arrS) {
    if (!checkNumberInRange(N)) {
        console.log("N\u306F".concat(MIN_VALUE, "\u4EE5\u4E0A").concat(MAX_VALUE, "\u4EE5\u4E0B\u306E\u6574\u6570\u3092\u5165\u529B\u3057\u3066\u304F\u3060\u3055\u3044"));
        return;
    }
    if (!checkLength(N, arrS)) {
        console.log("Nの値と配列の要素数が一致しません");
        return;
    }
    if (!checkArrayStringLengthInRange(arrS)) {
        console.log("\u6587\u5B57\u5217\u306F".concat(MIN_LENGTH, "\u4EE5\u4E0A").concat(MAX_LENGTH, "\u4EE5\u4E0B\u3067\u5165\u529B\u3057\u3066\u304F\u3060\u3055\u3044"));
        return;
    }
    // 配列の文字を結合する
    var joinStr = arrS.join("");
    console.log(joinStr.length);
};
calculateTotalStringLength(3, ["hello", "algo", "method"]);
calculateTotalStringLength(2, ["hello", "algo", "method"]); //エラーケース
calculateTotalStringLength(1001, ["hello", "algo", "method"]); //エラーケース
calculateTotalStringLength(1, ["a".repeat(20001)]);
