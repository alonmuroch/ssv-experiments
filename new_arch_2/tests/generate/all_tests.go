package main

import // AllTests maps package name to tests
(
	"ssv-experiments/new_arch_2/tests"
	asgard_qbft_full_flow "ssv-experiments/new_arch_2/tests/spec/asgard/qbft/full_flow"
	asgard_ssv_full_flow "ssv-experiments/new_arch_2/tests/spec/asgard/ssv/full_flow"
)

var AllTests = map[string][]tests.TestObject{"asgard_qbft_full_flow": asgard_qbft_full_flow.AllTests, "asgard_ssv_full_flow": asgard_ssv_full_flow.AllTests}
