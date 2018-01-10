package resources_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rcw5/pipeline-validator/resources"
	"github.com/rcw5/pipeline-validator/testhelpers"
)

var _ = Describe("Pipeline", func() {
	Context("Validate", func() {
		Context("Vars surrounded by brackets", func() {
			It("Validates a pipeline when all vars are provided", func() {
				p := resources.NewPipeline(testhelpers.SAMPLE_PIPELINE)
				results := p.Validate([]string{"var1", "var2"})
				Expect(results.IsSuccessful()).To(BeTrue())
			})
			It("Validates a pipeline when a var is missing", func() {
				p := resources.NewPipeline(testhelpers.SAMPLE_PIPELINE)
				results := p.Validate([]string{"var1"})
				Expect(results.MissingVarsError).To(MatchError("The following vars were present in the pipeline but not in the vars file: var2"))
			})
			It("Validates a pipeline when extra vars are present", func() {
				p := resources.NewPipeline(testhelpers.SAMPLE_PIPELINE)
				results := p.Validate([]string{"var1", "var2", "var3"})

				Expect(results.ExtraVarsError).To(MatchError("The following vars were present in the vars file but not the pipeline: var3"))
			})
			It("Validates a pipeline when both vars are missing and extra vars are present", func() {
				p := resources.NewPipeline(testhelpers.SAMPLE_PIPELINE)
				results := p.Validate([]string{"var1", "var3"})
				Expect(results.ExtraVarsError).To(MatchError("The following vars were present in the vars file but not the pipeline: var3"))
				Expect(results.MissingVarsError).To(MatchError("The following vars were present in the pipeline but not in the vars file: var2"))
			})
		})

		Context("Sorting the output", func() {
			It("Sorts the list of extra vars alphabetically", func() {
				p := resources.NewPipeline(testhelpers.SAMPLE_PIPELINE)
				results := p.Validate([]string{"var1", "var2", "zvar", "bvar"})
				Expect(results.MissingVarsError).ToNot(HaveOccurred())
				Expect(results.ExtraVarsError).To(MatchError("The following vars were present in the vars file but not the pipeline: bvar, zvar"))
			})
			It("Sorts the list of missing vars alphabetically", func() {
				p := resources.NewPipeline(testhelpers.SAMPLE_PIPELINE)
				results := p.Validate([]string{})
				Expect(results.ExtraVarsError).ToNot(HaveOccurred())
				Expect(results.MissingVarsError).To(MatchError("The following vars were present in the pipeline but not in the vars file: var1, var2"))
			})
		})

		Context("Vars surrounded by a mixture of curlys and brackets", func() {
			It("Validates a pipeline when all vars are provided", func() {
				p := resources.NewPipeline(testhelpers.SAMPLE_PIPELINE_MIXED)
				results := p.Validate([]string{"var1", "var2"})
				Expect(results.IsSuccessful()).To(BeTrue())
			})
			It("Validates a pipeline when a var is missing", func() {
				p := resources.NewPipeline(testhelpers.SAMPLE_PIPELINE_MIXED)
				results := p.Validate([]string{"var1"})
				Expect(results.MissingVarsError).To(MatchError("The following vars were present in the pipeline but not in the vars file: var2"))
			})
			It("Validates a pipeline when extra vars are present", func() {
				p := resources.NewPipeline(testhelpers.SAMPLE_PIPELINE_MIXED)
				results := p.Validate([]string{"var1", "var2", "var3"})

				Expect(results.ExtraVarsError).To(MatchError("The following vars were present in the vars file but not the pipeline: var3"))
			})
			It("Validates a pipeline when both vars are missing and extra vars are present", func() {
				p := resources.NewPipeline(testhelpers.SAMPLE_PIPELINE_MIXED)
				results := p.Validate([]string{"var1", "var3"})
				Expect(results.ExtraVarsError).To(MatchError("The following vars were present in the vars file but not the pipeline: var3"))
				Expect(results.MissingVarsError).To(MatchError("The following vars were present in the pipeline but not in the vars file: var2"))
			})
		})

		Context("Vars surrounded by a mixture of curlys only", func() {
			It("Validates a pipeline when all vars are provided", func() {
				p := resources.NewPipeline(testhelpers.SAMPLE_PIPELINE_CURLYS)
				results := p.Validate([]string{"var1", "var2"})
				Expect(results.IsSuccessful()).To(BeTrue())
			})
			It("Validates a pipeline when a var is missing", func() {
				p := resources.NewPipeline(testhelpers.SAMPLE_PIPELINE_CURLYS)
				results := p.Validate([]string{"var1"})
				Expect(results.MissingVarsError).To(MatchError("The following vars were present in the pipeline but not in the vars file: var2"))
			})
			It("Validates a pipeline when extra vars are present", func() {
				p := resources.NewPipeline(testhelpers.SAMPLE_PIPELINE_CURLYS)
				results := p.Validate([]string{"var1", "var2", "var3"})

				Expect(results.ExtraVarsError).To(MatchError("The following vars were present in the vars file but not the pipeline: var3"))
			})
			It("Validates a pipeline when both vars are missing and extra vars are present", func() {
				p := resources.NewPipeline(testhelpers.SAMPLE_PIPELINE_CURLYS)
				results := p.Validate([]string{"var1", "var3"})
				Expect(results.ExtraVarsError).To(MatchError("The following vars were present in the vars file but not the pipeline: var3"))
				Expect(results.MissingVarsError).To(MatchError("The following vars were present in the pipeline but not in the vars file: var2"))
			})
		})
	})
})
