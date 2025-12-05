package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println(minimumFlips(7))
}

func minimumFlips(n int) int {
	res := 0
	// convert int to binary
	binStr := ""
	for n > 0 {
		binStr = strconv.Itoa(n%2) + binStr
		n /= 2
	}
	//fmt.Println(binStr)
	// count minimum flip
	for i := 0; i < len(binStr)/2; i++ {
		if binStr[i] != binStr[len(binStr)-1-i] {
			res += 2
		}
	}
	return res
}

func generateTag(caption string) string {
	arr := strings.Split(caption, " ")
	res := "#"
	for _, s := range arr {
		s = strings.TrimSpace(s)
		if len(s) == 0 {
			continue
		}
		if len(res) == 1 {
			res += strings.ToLower(s)
			continue
		}

		res += strings.ToUpper(s[0:1]) + strings.ToLower(s[1:])
	}

	if len(res) > 100 {
		return res[:100]
	}

	return res

	// res := "#"
	// isSpace := false
	// for _, s := range caption {
	// 	if len(res) == 1 && s != ' ' {
	// 		res += strings.ToLower(string(s))
	// 		isSpace = false
	// 		continue
	// 	}

	// 	if s == ' ' {
	// 		isSpace = true
	// 		continue
	// 	}

	// 	if isSpace {
	// 		res += strings.ToUpper(string(s))
	// 		isSpace = false
	// 		continue
	// 	}

	// 	if s >= 'A' && s <= 'Z' {
	// 		res += strings.ToLower(string(s))
	// 		continue
	// 	}

	// 	res += string(s)
	// }

	// if len(res) > 100 {
	// 	return res[:100]
	// }
	// // fmt.Println("len data: ", len(res))
	// // fmt.Println("data: ", res)
	// return res
}

// Example 1:
// Input: nums1 = [4,1,2], nums2 = [1,3,4,2]
// Output: [-1,3,-1]
// Explanation: The next greater element for each value of nums1 is as follows:
// - 4 is underlined in nums2 = [1,3,4,2]. There is no next greater element, so the answer is -1.
// - 1 is underlined in nums2 = [1,3,4,2]. The next greater element is 3.
// - 2 is underlined in nums2 = [1,3,4,2]. There is no next greater element, so the answer is -1.
//
// Example 2:
// Input: nums1 = [2,4], nums2 = [1,2,3,4]
// Output: [3,-1]
// Explanation: The next greater element for each value of nums1 is as follows:
// - 2 is underlined in nums2 = [1,2,3,4]. The next greater element is 3.
// - 4 is underlined in nums2 = [1,2,3,4]. There is no next greater element, so the answer is -1.
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	mapIntNext := make(map[int]int)
	stack := []int{}
	for _, n := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < n {
			mapIntNext[stack[len(stack)-1]] = n
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, n)
	}

	for _, n := range nums1 {
		if v, ok := mapIntNext[n]; ok {
			res = append(res, v)
		} else {
			res = append(res, -1)
		}
	}
	return res
}

type RecentCounter struct {
	Queue []int
}

func Constructor() RecentCounter {
	queue := make([]int, 0)
	return RecentCounter{
		Queue: queue,
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.Queue = append(this.Queue, t)
	for len(this.Queue) > 0 && this.Queue[0] < t-3000 {
		this.Queue = this.Queue[1:]
	}
	return len(this.Queue)
}

//You have a RecentCounter class which counts the number of recent requests within a certain time frame.
//
//Implement the RecentCounter class:
//
//RecentCounter() Initializes the counter with zero recent requests.
//int ping(int t) Adds a new request at time t, where t represents some time in milliseconds, and returns the number of requests that has happened in the past 3000 milliseconds (including the new request). Specifically, return the number of requests that have happened in the inclusive range [t - 3000, t].
//It is guaranteed that every call to ping uses a strictly larger value of t than the previous call.

func decodeString(s string) string {
	stackNum := []int{}
	stackStr := []string{}
	curStr := ""
	curNum := 0

	for i := 0; i < len(s); i++ {
		ch := s[i]

		if ch >= '0' && ch <= '9' {
			// Xây dựng số khi gặp chữ số
			curNum = curNum*10 + int(ch-'0')
		} else if ch == '[' {
			stackNum = append(stackNum, curNum)
			stackStr = append(stackStr, curStr)
			curNum = 0
			curStr = ""
		} else if ch == ']' {
			num := stackNum[len(stackNum)-1]
			preStr := stackStr[len(stackStr)-1]
			repeatedStr := ""
			for j := 0; j < num; j++ {
				repeatedStr += curStr
			}
			curStr = preStr + repeatedStr
			stackNum = stackNum[:len(stackNum)-1]
			stackStr = stackStr[:len(stackStr)-1]
		} else {
			curStr += string(ch)
		}
	}

	return curStr
}

func asteroidCollision(asteroids []int) []int {
	stackAsteroids := make([]int, 0)
	for _, a := range asteroids {
		if a > 0 {
			stackAsteroids = append(stackAsteroids, a)
		} else {
			for len(stackAsteroids) > 0 && stackAsteroids[len(stackAsteroids)-1] > 0 && getAbs(a) > getAbs(stackAsteroids[len(stackAsteroids)-1]) {
				stackAsteroids = stackAsteroids[:len(stackAsteroids)-1]
			}

			if len(stackAsteroids) == 0 || stackAsteroids[len(stackAsteroids)-1] < 0 {
				stackAsteroids = append(stackAsteroids, a)
			} else if getAbs(a) == getAbs(stackAsteroids[len(stackAsteroids)-1]) {
				stackAsteroids = stackAsteroids[:len(stackAsteroids)-1]
			}
		}
	}
	return stackAsteroids
}

func getAbs(a int) int {
	if a >= 0 {
		return a
	}

	return -a
}

func removeStars(s string) string {
	res := ""
	cntStars := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '*' {
			j := i - 1
			cntStars++

			for j >= 0 {
				if s[j] == '*' {
					cntStars++
					j--
					continue
				} else {
					if cntStars == 0 {
						break
					}

					if cntStars > 0 {
						cntStars--
						j--
						continue
					}
				}
			}

			i = j
			//fmt.Println("======: i :, ", i, " j: ", j)
		}

		if i >= 0 {
			res = string(s[i]) + res
		}

	}
	return res
}

func equalPairs(grid [][]int) int {
	res := 0
	mapRow := make(map[string]int)
	for _, row := range grid {
		strRow := ""
		for _, col := range row {
			strRow += strconv.Itoa(col) + ","
		}
		mapRow[strRow]++
	}

	mapCol := make(map[string]int)
	for col, _ := range grid {
		strCol := ""
		for row := 0; row < len(grid); row++ {
			strCol += strconv.Itoa(grid[row][col]) + ","
		}
		mapCol[strCol]++
	}

	for k, vRow := range mapRow {
		if vCol, ok := mapCol[k]; ok {
			res += vRow * vCol
		}
	}
	return res
}
func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	mapChar1 := make(map[byte]int)
	mapChar2 := make(map[byte]int)
	for i := 0; i < len(word1); i++ {
		mapChar1[word1[i]]++
		mapChar2[word2[i]]++
	}

	for k, _ := range mapChar1 {
		if mapChar2[k] == 0 {
			return false
		}
	}

	mapFreq1 := make(map[int]int)
	mapFreq2 := make(map[int]int)
	for _, v := range mapChar1 {
		mapFreq1[v]++
	}
	for _, v := range mapChar2 {
		mapFreq2[v]++
	}

	for k, v := range mapFreq1 {
		if mapFreq2[k] != v {
			return false
		}
	}
	return true
}

func longestSubarray(nums []int) int {
	l, r := 0, 0
	maxLen, cntZero := 0, 0
	for r < len(nums) {
		if nums[r] == 0 {
			cntZero++
		}

		if cntZero > 1 {
			if nums[l] == 0 {
				cntZero--
			}
			l++
		}

		if r-l+1 > maxLen {
			maxLen = r - l + 1
		}
		r++
	}

	return maxLen - 1
}

func strongPasswordChecker1(password string) int {
	n := len(password)
	hasLower, hasUpper, hasDigit := 0, 0, 0

	for _, ch := range password {
		if 'a' <= ch && ch <= 'z' {
			hasLower = 1
		} else if 'A' <= ch && ch <= 'Z' {
			hasUpper = 1
		} else if '0' <= ch && ch <= '9' {
			hasDigit = 1
		}
	}
	missingTypes := 3 - (hasLower + hasUpper + hasDigit)

	repeats := []int{}
	for i := 0; i < n; {
		j := i
		for j < n && password[j] == password[i] {
			j++
		}
		length := j - i
		if length >= 3 {
			repeats = append(repeats, length)
		}
		i = j
	}

	if n < 6 {
		return max(missingTypes, 6-n)
	}

	if n <= 20 {
		replace := 0
		for _, length := range repeats {
			replace += length / 3
		}
		return max(missingTypes, replace)
	}

	deleteCount := n - 20
	leftDelete := deleteCount

	mod0, mod1, mod2 := []int{}, []int{}, []int{}
	for _, length := range repeats {
		if length%3 == 0 {
			mod0 = append(mod0, length)
		} else if length%3 == 1 {
			mod1 = append(mod1, length)
		} else {
			mod2 = append(mod2, length)
		}
	}

	for i := 0; i < len(mod0) && leftDelete > 0; i++ {
		need := 1
		if leftDelete >= need {
			mod0[i] -= need
			leftDelete -= need
			if mod0[i] < 3 {
				mod0[i] = 0
			}
		}
	}

	for i := 0; i < len(mod1) && leftDelete > 1; i++ {
		need := 2
		if leftDelete >= need {
			mod1[i] -= need
			leftDelete -= need
			if mod1[i] < 3 {
				mod1[i] = 0
			}
		}
	}

	for i := 0; i < len(mod2) && leftDelete > 2; i++ {
		need := 3
		if leftDelete >= need {
			mod2[i] -= need
			leftDelete -= need
			if mod2[i] < 3 {
				mod2[i] = 0
			}
		}
	}

	replace := 0
	for _, length := range mod0 {
		if length >= 3 {
			replace += length / 3
		}
	}
	for _, length := range mod1 {
		if length >= 3 {
			replace += length / 3
		}
	}
	for _, length := range mod2 {
		if length >= 3 {
			replace += length / 3
		}
	}

	return deleteCount + max(missingTypes, replace)
}

func strongPasswordChecker(password string) int {
	// len < 6 => insert cho du 6
	if len(password) < 6 {
		return 6 - len(password)
	}
	cntDelete := len(password) - 20
	isNumber, isLower, isUpper, isChecked := false, false, false, false
	cntRepeat, cntChar, cntConditionChar := 0, 1, 3 // cntConditionChar = 3 la do dk chu thuong, chu hoa, so
	charRepeat := string(password[0])
	for i := 1; i < len(password); i++ {
		// tim so lan repeat can sua
		if string(password[i]) == charRepeat {
			cntChar++
			if len(password) >= 20 && cntChar > 3 && cntChar%3 > 0 && cntChar%3 < 3 {
				if !isChecked {
					isChecked = true
					cntRepeat--
				}
			} else if cntChar%3 == 0 {
				cntRepeat++
			}
		} else {
			isChecked = false
			charRepeat = ""
			cntChar = 0
			charRepeat = string(password[i])
			cntChar = 1
		}

		if password[i] >= 48 && password[i] <= 57 && !isNumber {
			isNumber = true
			cntConditionChar--
		}

		if password[i] >= 65 && password[i] <= 90 && !isUpper {
			isUpper = true
			cntConditionChar--
		}

		if password[i] >= 97 && password[i] <= 122 && !isLower {
			isLower = true
			cntConditionChar--
		}
	}

	if len(password) >= 6 && len(password) <= 20 {
		return max(cntRepeat, cntConditionChar)
	}

	return cntDelete + max(cntRepeat, cntConditionChar)
}

func longestOnes(nums []int, k int) int {
	l, r := 0, 0
	cntZero := 0
	maxLen := 0
	for r < len(nums) {
		// neu r la 0 thi lat => tang bien dem
		if nums[r] == 0 {
			cntZero++
		}

		// neu bien dem lown hon so luong cho phep thi phai tang left pointer va giam bien dem
		for cntZero > k {
			if nums[l] == 0 {
				cntZero--
			}
			l++
		}

		//check max length
		if r-l+1 > maxLen {
			maxLen = r - l + 1
		}
		r++
	}

	return maxLen
}
func replaceElements(arr []int) []int {
	max := -1
	res := make([]int, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		res = append([]int{max}, res...)

		if max < arr[i] {
			max = arr[i]
		}
	}
	return res
}

func dayOfTheWeek(day int, month int, year int) string {
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	// fmt.Println("====: ", t.Weekday())
	daysOfWeek := [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	return string(daysOfWeek[t.Weekday()])
}
func alternateDigitSum(n int) int {
	str := strconv.Itoa(n)
	sum := 0
	for i := 0; i < len(str); i++ {
		num, _ := strconv.Atoi(string(str[i]))
		if i%2 == 0 {
			sum += num
		} else {
			sum -= num
		}
	}
	return sum
}
func isUgly(n int) bool {
	if n < 0 {
		n *= -1
	}
	for n > 1 {
		if n%2 == 0 {
			n /= 2
			continue
		} else if n%3 == 0 {
			n /= 3
			continue
		} else if n%5 == 0 {
			n /= 5
			continue
		}

		return false
	}
	return true
}

func convertToTitle(columnNumber int) string {
	bytes := make([]byte, 0)
	for columnNumber > 0 {
		columnNumber--                                   // trừ đi 1 để  đảm bảo số dư từ 0-25, trong khi các cột excel từ 1-26
		bytes = append(bytes, 'A'+byte(columnNumber%26)) // ('A' + byte(columnNumber%26))
		columnNumber /= 26
	}
	for i := 0; i < len(bytes)/2; i++ {
		bytes[i], bytes[len(bytes)-i-1] = bytes[len(bytes)-i-1], bytes[i]
	}
	return string(bytes) //result.WriteString(string(result.Bytes()))
}
func titleToNumber(columnTitle string) int {
	sum := float64(0)
	mapAlpha := map[byte]int{
		'A': 1, 'B': 2, 'C': 3, 'D': 4, 'E': 5, 'F': 6, 'G': 7, 'H': 8, 'I': 9, 'J': 10, 'K': 11, 'L': 12, 'M': 13,
		'N': 14, 'O': 15, 'P': 16, 'Q': 17, 'R': 18, 'S': 19, 'T': 20, 'U': 21, 'V': 22, 'W': 23, 'X': 24, 'Y': 25, 'Z': 26,
	} //"ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := len(columnTitle) - 1; i >= 0; i-- {
		sum += float64(mapAlpha[columnTitle[i]]) * math.Pow(float64(26), float64(len(columnTitle)-1-i))
	}
	return int(sum)
}

func toGoatLatin(sentence string) string {
	res := ""
	mapVowel := map[byte]bool{'u': true, 'e': true, 'o': true, 'i': true, 'a': true,
		'U': true, 'E': true, 'O': true, 'I': true, 'A': true}
	words := strings.Split(sentence, " ")
	for i, w := range words {
		if mapVowel[w[0]] {
			res += w + "ma"

		} else {
			res += w[1:] + string(w[0]) + "ma"
		}
		// fmt.Println("adasd: ", i, res)
		for j := 0; j <= i; j++ {
			res += "a"
		}
		// fmt.Println("adasd111: ", res)
		if i == len(words)-1 {
			break
		}
		res += " "
	}
	return res
}

func modifyString(s string) string {
	bytes := []byte(s)
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == '?' {
			if i == 0 {
				if i+1 < len(s) {
					bytes[i] = getDiffChar(' ', bytes[i+1])
				} else {
					bytes[i] = getDiffChar(' ', bytes[i])
				}

			} else if i == len(s)-1 {
				bytes[i] = getDiffChar(' ', bytes[i-1])
			} else {
				bytes[i] = getDiffChar(bytes[i-1], bytes[i+1])
			}
		}
	}
	return string(bytes)
}
func getDiffChar(c1 byte, c2 byte) byte {
	if c1 != 'a' && c2 != 'a' {
		return 'a'
	} else if c1 != 'b' && c2 != 'b' {
		return 'b'
	} else {
		return 'c'
	}
}
func maxPower(s string) int {
	cnt, max := 1, math.MinInt32
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cnt++
		} else {
			if cnt > max {
				max = cnt
			}

			cnt = 1
		}
	}

	if cnt > max {
		return cnt
	}
	return max
}

func daysBetweenDates(date1 string, date2 string) int {
	// Chuyển đổi chuỗi thành đối tượng time
	layout := "2006-01-02" // Định dạng ngày tháng
	t1, _ := time.Parse(layout, date1)
	t2, _ := time.Parse(layout, date2)

	// Tính hiệu số giữa hai thời điểm và lấy giá trị tuyệt đối
	diff := t2.Sub(t1)
	return int(diff.Hours() / 24) // Chuyển đổi từ giờ sang ngày
}
func freqAlphabets(s string) string {
	alpha := "abcdefghijklmnopqrstuvwxyz"
	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != '#' {
			idx, _ := strconv.Atoi(string(s[i]))
			// res = string(s[i]) + res
			res = string(alpha[idx-1]) + res
		} else {
			idx, _ := strconv.Atoi(s[i-2 : i])
			res = string(alpha[idx-1]) + res

			// reset i
			i -= 2
		}
	}
	return res
}

func largeGroupPositions(s string) [][]int {
	res := make([][]int, 0)
	// l, r := 0, 0
	start, count := 0, 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			if count > 2 {
				res = append(res, []int{start, start + count - 1})
			}

			start = i
			count = 1
		}
	}
	// đoạn cuối, nếu count >= 3 thì sẽ phải append thêm dãy cuối
	if count > 2 {
		res = append(res, []int{start, start + count - 1})
	}
	fmt.Println("data: ", start, count)
	return res
}

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	// 1 <= nums1.length, nums2.length <= 200
	// nums1[i].length == nums2[j].length == 2
	// 1 <= idi, vali <= 1000
	// Both arrays contain unique ids.
	// Both arrays are in strictly ascending order by id.

	i1, i2, max := 0, 0, 0
	mapTotal := make(map[int]int)
	for i1 < len(nums1) || i2 < len(nums2) {
		if i1 < len(nums1) {
			mapTotal[nums1[i1][0]] += nums1[i1][1]
			if nums1[i1][0] > max {
				max = nums1[i1][0]
			}

			i1++
		}

		if i2 < len(nums2) {
			mapTotal[nums2[i2][0]] += nums2[i2][1]
			if nums2[i2][0] > max {
				max = nums2[i2][0]
			}

			i2++
		}
	}

	res := [][]int{}
	for i := 1; i <= max; i++ {
		if val, ok := mapTotal[i]; ok {
			res = append(res, []int{i, val})
		}
	}

	return res
}
func makeSmallestPalindrome(s string) string {
	l, r := 0, len(s)-1
	bytes := []byte(s)
	for l < r {
		if bytes[l] > bytes[r] {
			bytes[l] = bytes[r]
		} else if bytes[l] < bytes[r] {
			bytes[r] = bytes[l]
		}
		l++
		r--
	}
	return string(bytes)
}
func reorderSpaces(text string) string {
	res := ""
	cntSpace := 0
	strs := make([]string, 0)
	w := ""
	for _, c := range text {
		if c == ' ' {
			cntSpace++
			if len(w) > 0 {
				strs = append(strs, w)
			}
			w = ""
			continue
		}

		w += string(c)

	}
	strSpace := ""

	i := 0
	for i < cntSpace/(len(strs)-1) {
		strSpace += " "
		i++
	}
	res = strings.Join(strs, strSpace)

	for len(res) < len(text) {
		res += " "
	}
	fmt.Println("asdasdaD: ", len(res), len(text))
	return res
}

func cellsInRange(s string) []string {
	cStart, cEnd := s[0], s[3]
	iStart, iEnd := s[1], s[4]
	res := make([]string, 0)
	for cStart <= cEnd {
		for iStart <= iEnd {
			res = append(res, string(cStart)+string(iStart))

			iStart++
		}
		cStart++
		iStart = s[1]
	}

	fmt.Println("AHHHHHH: ", res)
	return res
}

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	maxStart, minEnd := countDayPerson(arriveAlice, leaveAlice)
	BobStart, BobEnd := countDayPerson(arriveBob, leaveBob)
	if BobEnd < minEnd {
		minEnd = BobEnd
	}

	if maxStart < BobStart {
		maxStart = BobStart
	}
	if maxStart < minEnd {
		return minEnd - maxStart + 1
	}
	return 0
}

func countDayPerson(arrive string, leave string) (int, int) {
	dayStart, dayEnd := 0, 0
	dayInMonths := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	monthA, dayA := arrive[:2], arrive[3:]
	mA, _ := strconv.Atoi(monthA)
	dA, _ := strconv.Atoi(dayA)
	for m, d := range dayInMonths {
		if m+1 == mA {
			dayStart += dA
			break
		}

		dayStart += d
	}

	monthL, dayL := leave[:2], leave[3:]
	mL, _ := strconv.Atoi(monthL)
	dL, _ := strconv.Atoi(dayL)

	for m, d := range dayInMonths {
		if m+1 == mL {
			dayEnd += dL
			break
		}

		dayEnd += d
	}
	fmt.Println("day: ", dayStart, dayEnd)
	return dayStart, dayEnd
}

func isBalanced(num string) bool {
	sum, cnt := 0, 0
	for i, c := range num {
		cnt++
		if i%2 == 0 {
			sum += int(c)
			continue
		}
		sum -= int(c)
	}
	fmt.Println("count: ", cnt, sum)
	return sum == 48 || sum == 0
}

func countValidWords(sentence string) int {
	res := 0
	words := strings.Split(sentence, " ")
	for _, w := range words {
		if isWordValid(w) && len(w) > 0 {
			res++
		}
	}
	return res
}

func isWordValid(s string) bool {
	cntHyphen, cntMark := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] <= 57 && s[i] >= 48 {
			return false
		}

		// dấu câu phải nằm cuối cùng của từ
		if (s[i] < 97 || s[i] > 122) && s[i] != '-' {
			cntMark++
			if cntMark > 1 {
				return false
			}
			if i != len(s)-1 {
				return false
			}

		}

		// nếu là dấu gạch phải nằm giữa 2 chữ cái
		if s[i] == '-' {
			cntHyphen++
			if cntHyphen > 1 {
				return false
			}
			if i == 0 || i == len(s)-1 {
				return false
			}

			if s[i+1] < 97 || s[i+1] > 122 || s[i-1] < 97 || s[i-1] > 122 {
				return false
			}
		}
	}
	return true
}
func areNumbersAscending(s string) bool {
	strs := strings.Split(s, " ")
	pre := 0
	for _, str := range strs {
		number, err := strconv.Atoi(str)
		if err != nil {
			continue
		}
		if number > pre {
			pre = number
		} else {
			return false
		}
	}
	return true
}

// "1 box has 3 blue 4 red 6 green and 12 yellow marbles"
func mostFrequent(nums []int, key int) int {
	// Tạo một map để đếm tần suất
	freq := make(map[int]int)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == key {
			freq[nums[i+1]]++
		}
	}

	// Tìm số có tần suất lớn nhất
	maxFreq := 0
	result := 0
	for num, count := range freq {
		if count > maxFreq {
			maxFreq = count
			result = num
		}
	}

	return result
}
func decodeMessage(key string, message string) string {
	alpha := "abcdefghijklmnopqrstuvwxyz"
	mapChar := make(map[string]string)
	iKey, iAlpha := 0, 0
	// the quick brown fox jumps over the lazy dog
	// abc defgh ijklm n o p qrs  t       uvwx y z
	for iAlpha < len(alpha) && iKey < len(key) {
		// là chữ cái và chưa có trong map mới thêm vào map
		if _, ok := mapChar[string(key[iKey])]; !ok && key[iKey] >= 97 && key[iKey] <= 122 {
			mapChar[string(key[iKey])] = string(alpha[iAlpha])
			// sau khi thêm vào map thì tăng index lên 1
			iKey++
			iAlpha++
		} else {
			// nếu chữ cái của key đã có trong map thì next sang chữ tiếp theo
			// nếu ký tự trong key là khoảng trắng thì next sang ký tự tiếp theo
			iKey++
		}
	}
	res := ""
	for i := 0; i < len(message); i++ {
		if message[i] == ' ' {
			res += " "
		}
		res += mapChar[string(message[i])]
	}
	fmt.Println(mapChar)
	return res
}

func digitCount(num string) bool {
	// 1210 map[0->1, 1->2, 2->1]
	// 0123
	mapDigit := make(map[byte]int)
	for i := 0; i < len(num); i++ {
		mapDigit[num[i]]++
	}
	fmt.Println("data map: ", mapDigit)
	for i := 0; i < len(num); i++ {
		fmt.Println("====: ", int(num[i]), byte(i+48), mapDigit[byte(i+48)])
		if (int(num[i]) - 48) != mapDigit[byte(i+48)] {
			return false
		}
	}
	return true
}

func reverseStr(s string, k int) string {
	/*
		adcb defg hijk lmn
		0-3  4-7  8-11 12-14
		end = i+k-1 (i+=k)
		i=0, end=3
		i=4, end=7
		i=8, end=11
		i=12, end=15
		end > len(s) - 1 => end = len(s)-1=14
	*/
	return ""
}
func reverseString(s []byte) {

	fmt.Println("====BEFORE: ", s)
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
	fmt.Println("====AFTER: ", s)
}
func isHappy(n int) bool {
	mapNum := make(map[int]bool)
	sum, tmp := 0, n

	for {
		sum = 0
		for tmp > 0 {
			sum += (tmp % 10) * (tmp % 10)
			tmp /= 10

		}
		fmt.Println("=sum: ", sum)
		if mapNum[sum] {
			return false
		}
		if sum == 1 {
			return true
		}

		mapNum[sum] = true
		tmp = sum
	}

}
func isPalindrome(s string) bool {
	// c1 => tao 1 str chuẩn rồi đi từ 2 đầu thôi
	// strFormat := ""
	// for i := 0; i < len(s); i++ {
	// 	// chu thuong va so
	// 	if (s[i] >= 97 && s[i] <= 122) || (s[i] >= 48 && s[i] <= 57) {
	// 		strFormat += string(s[i])
	// 	} else if s[i] >= 65 && s[i] <= 90 {
	// 		// chu hoa
	// 		strFormat += string(s[i] + 32)
	// 	}
	// }
	// l, r := 0, len(strFormat)-1
	// for l < r {
	// 	if strFormat[l] != strFormat[r] {
	// 		return false
	// 	}
	// 	l++
	// 	r--
	// }
	// fmt.Println("=data: ", strFormat)
	// return true

	// c2 2 con trỏ đi đên đâu check đến đó
	l, r := 0, len(s)-1
	for l < r {
		s = strings.ToLower(s)
		if !((s[l] >= 97 && s[l] <= 122) || (s[l] >= 48 && s[l] <= 57)) {
			l++
			continue
		}

		if !((s[r] >= 97 && s[r] <= 122) || (s[r] >= 48 && s[r] <= 57)) {
			r--
			continue
		}

		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}
	return true
}
func charValid(c byte) bool {
	return (c >= 97 && c <= 122) || (c >= 48 && c <= 57)
}

func destCity(paths [][]string) string {
	freq := make(map[string]int)
	mapPos := make(map[string]int)
	for _, p := range paths {
		freq[p[0]]++
		freq[p[1]]++
		mapPos[p[0]] = 0
		mapPos[p[1]] = 1
	}
	res := ""
	for k, v := range freq {
		if v == 1 && mapPos[k] == 1 {
			res = k
		}
	}
	return res
}

func countLargestGroup(n int) int {
	mapDigitToSum := make(map[int][]int)
	maxLen := -1
	for i := 1; i <= n; i++ {
		sumDigit := sumDigit(i)
		if _, ok := mapDigitToSum[sumDigit]; !ok {
			mapDigitToSum[sumDigit] = []int{i}

		} else {
			mapDigitToSum[sumDigit] = append(mapDigitToSum[sumDigit], i)
		}

		if len(mapDigitToSum[sumDigit]) > maxLen {
			maxLen = len(mapDigitToSum[sumDigit])
		}

	}
	res := 0
	for _, v := range mapDigitToSum {
		if len(v) == maxLen {
			res++
		}
	}
	return res
}

func sumDigit(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
func countVowelSubstrings(word string) int {
	mapVowel := map[byte]bool{'u': true, 'e': true, 'o': true, 'a': true, 'i': true}

	res := 0
	for i := 0; i < len(word); i++ {
		str := ""
		for j := i; j < len(word); j++ {
			if mapVowel[word[j]] {
				str += string(word[j])
			} else {
				str = ""
			}

			if len(str) > 4 {
				mapCheck := make(map[byte]bool)
				for k := 0; k < len(str); k++ {
					mapCheck[str[k]] = true
				}

				if len(mapCheck) == len(mapVowel) {
					res++
				}
			}
		}
		fmt.Println("data: ", str)

	}

	return res
}
