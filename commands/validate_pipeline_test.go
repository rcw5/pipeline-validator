package commands_test

import (
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rcw5/pipeline-validator/commands"
	"github.com/rcw5/pipeline-validator/testhelpers"
)

var _ = Describe("Validate Pipeline", func() {

	var tempDir string
	var pipelinePath string
	BeforeEach(func() {
		var err error
		tempDir, err = ioutil.TempDir("", "vars-validator")
		Expect(err).ToNot(HaveOccurred())
		pipelinePath, err = testhelpers.WriteStringToTempFile(tempDir, testhelpers.SAMPLE_PIPELINE)
		Expect(err).ToNot(HaveOccurred())
	})
	AfterEach(func() {
		os.Remove(tempDir)
	})
	It("Validates a pipeline where all vars are provided", func() {
		varsPath, err := testhelpers.WriteStringToTempFile(tempDir, testhelpers.SAMPLE_VARS)
		Expect(err).ToNot(HaveOccurred())
		err = commands.ValidatePipeline(pipelinePath, varsPath)
		Expect(err).ToNot(HaveOccurred())
	})

	It("Handles vars spread across multiple files", func() {
		varsPath1, err := testhelpers.WriteStringToTempFile(tempDir, testhelpers.SAMPLE_VARS_PART1)
		Expect(err).ToNot(HaveOccurred())
		varsPath2, err := testhelpers.WriteStringToTempFile(tempDir, testhelpers.SAMPLE_VARS_PART2)
		Expect(err).ToNot(HaveOccurred())

		err = commands.ValidatePipeline(pipelinePath, varsPath1, varsPath2)
		Expect(err).ToNot(HaveOccurred())
	})

	It("Handles pipelines with a mixture of curly braces and brackets", func() {
		pipelinePath, err := testhelpers.WriteStringToTempFile(tempDir, testhelpers.SAMPLE_PIPELINE_MIXED)
		Expect(err).ToNot(HaveOccurred())
		varsPath, err := testhelpers.WriteStringToTempFile(tempDir, testhelpers.SAMPLE_VARS)
		Expect(err).ToNot(HaveOccurred())

		err = commands.ValidatePipeline(pipelinePath, varsPath)
		Expect(err).ToNot(HaveOccurred())
	})

	Context("Errors", func() {
		It("Returns a list of vars in the pipeline but not declared in vars files", func() {
			varsPath, err := testhelpers.WriteStringToTempFile(tempDir, testhelpers.SAMPLE_MISSING_VARS)
			Expect(err).ToNot(HaveOccurred())
			err = commands.ValidatePipeline(pipelinePath, varsPath)
			Expect(err).To(MatchError("The following vars were present in the pipeline but not in the vars file: var2"))
		})
		It("Returns a list of vars that are in vars file but not used in the pipeline", func() {
			varsPath, err := testhelpers.WriteStringToTempFile(tempDir, testhelpers.SAMPLE_EXTRA_VARS)
			Expect(err).ToNot(HaveOccurred())
			err = commands.ValidatePipeline(pipelinePath, varsPath)
			Expect(err).To(MatchError("The following vars were present in the vars file but not the pipeline: var3"))
		})
	})
})
