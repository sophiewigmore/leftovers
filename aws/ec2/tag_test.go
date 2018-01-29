package ec2_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/genevievelesperance/leftovers/aws/ec2"
	"github.com/genevievelesperance/leftovers/aws/ec2/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tag", func() {
	var (
		tag        ec2.Tag
		client     *fakes.TagsClient
		key        *string
		value      *string
		resourceId *string
	)

	BeforeEach(func() {
		client = &fakes.TagsClient{}
		key = aws.String("the-key")
		value = aws.String("the-value")
		resourceId = aws.String("the-resource-id")

		tag = ec2.NewTag(client, key, value, resourceId)
	})

	It("releases the tag", func() {
		err := tag.Delete()
		Expect(err).NotTo(HaveOccurred())

		Expect(client.DeleteTagsCall.CallCount).To(Equal(1))
		Expect(client.DeleteTagsCall.Receives.Input.Tags).To(HaveLen(1))
		Expect(client.DeleteTagsCall.Receives.Input.Tags[0].Key).To(Equal(key))
		Expect(client.DeleteTagsCall.Receives.Input.Tags[0].Value).To(Equal(value))
		Expect(client.DeleteTagsCall.Receives.Input.Resources).To(HaveLen(1))
		Expect(client.DeleteTagsCall.Receives.Input.Resources[0]).To(Equal(resourceId))
	})

	Context("the client fails", func() {
		BeforeEach(func() {
			client.DeleteTagsCall.Returns.Error = errors.New("banana")
		})

		It("returns the error", func() {
			err := tag.Delete()
			Expect(err).To(MatchError("FAILED deleting tag the-value: banana"))
		})
	})
})