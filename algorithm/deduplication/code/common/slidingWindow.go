package main

import (
	"math"
)

/**
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

注意：

对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。
示例 1：

输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
示例 2：

输入：s = "a", t = "a"
输出："a"
解释：整个字符串 s 是最小覆盖子串。
示例 3:

输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。
*/

func minWindows(s string, t string) string {
	need, windows := make(map[byte]int), make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}

	left, right := 0, 0
	valid := 0
	start, length := 0, math.MaxInt
	for right < len(s) {
		w := s[right]
		if _, ok := need[w]; ok {
			windows[w]++
			if windows[w] == need[w] {
				valid++
			}
		}

		// 缩小窗口
		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}
			p := s[left]
			left++
			if _, ok := need[p]; ok {
				if windows[p] == need[p] {
					valid--
				}
				windows[p]--
			}
		}
	}

	if length == math.MaxInt {
		return ""
	}

	return s[start : start+length]

}

/**
567. 字符串的排列 | 力扣 | LeetCode |
给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。

换句话说，s1 的排列之一是 s2 的 子串 。

示例 1：

输入：s1 = "ab" s2 = "eidbaooo"
输出：true
解释：s2 包含 s1 的排列之一 ("ba").
示例 2：

输入：s1= "ab" s2 = "eidboaoo"
输出：false
提示：

1 <= s1.length, s2.length <= 104
s1 和 s2 仅包含小写字母
*/

func checkInclusion(s string, t string) bool {
	need, windows := make(map[rune]int), make(map[rune]int)
	for _, v := range s {
		need[v]++
	}
	left, right := 0, 0
	valid := 0

	for right < len(s) {
		w := s[right]
		if _, ok := need[rune(w)]; ok {
			windows[rune(w)]++
			if windows[rune(w)] == need[rune(w)] {
				valid++
			}
		}
		// 确认边界条件
		for right-left >= len(need) {
			if valid == len(need) {
				return true
			}
			p := s[left]
			left++
			if _, ok := need[rune(p)]; ok {
				if windows[rune(w)] == need[rune(w)] {
					valid--
				}
				windows[rune(p)]--
			}
		}

	}

	return false

}

/**
438. 找到字符串中所有字母异位词 | 力扣 | LeetCode |
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。

示例 1:

输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
 示例 2:

输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
提示:

1 <= s.length, p.length <= 3 * 104
s 和 p 仅包含小写字母
这个所谓的字母异位词，不就是排列吗，搞个高端的说法就能糊弄人了吗？相当于，输入一个串 S，一个串 T，找到 S 中所有 T 的排列，返回它们的起始索引。
*/
// findAnagrams
func findAnagrams(s string, t string) []int {
	need, windows := make(map[byte]int), make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}
	left, right := 0, 0
	valid := 0
	var res []int
	for right < len(s) {
		w := s[right]
		right++
		if _, ok := need[w]; ok {
			windows[w]++
			if windows[w] == need[w] {
				valid++
			}
		}
		for right-left >= len(need) {
			if valid == len(need) {
				res = append(res, left)
			}
			p := s[left]
			left++
			if _, ok := need[p]; ok {
				if windows[w] == need[w] {
					valid--
				}
				windows[p]--
			}
		}

	}
	return nil

}

/**
3. 无重复字符的最长子串 | 力扣 | LeetCode |
给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。

示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
提示：

0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成
*/

func lengthOFLongestSubstring(nums []int) int {
	left, right := 0, 0
	res := 0
	windows := make(map[byte]int)
	for right < len(nums) {
		w := nums[right]
		windows[byte(w)]++
		for windows[byte(w)] > 1 {
			p := nums[left]
			left++
			windows[byte(p)]--
		}
		// 更新长度
		if res < (right - left) {
			res = right - left
		}
	}

	return res
}
