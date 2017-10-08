package ping

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"

    "context"
    "net"
)

func TestPing(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Ping Suite")
}

type TestHost struct{
	DnsName  	string
	Ip4 			net.IP
	Ip6 			net.IP
}

var localhost = TestHost{
	"localhost",
	net.ParseIP("127.0.0.1"),
	net.ParseIP("::1"),
}

var _ = Describe("NewPinger", func() {
	It("returns pointer to a new Pinger", func() {
		pinger, _ := NewPinger(context.TODO(), localhost.DnsName)
		Expect(pinger).ToNot(BeNil())
		Expect(pinger).To(BeAssignableToTypeOf(&Pinger{}))
	})

	It("parses ip4 address", func() {
		pinger, _ := NewPinger(context.TODO(), localhost.Ip4.String())
		Expect(pinger.IPAddr().IP).To(Equal(localhost.Ip4))
	})

	It("parses ip6 address", func() {
		pinger, _ := NewPinger(context.TODO(), localhost.Ip6.String())
		Expect(pinger.IPAddr().IP).To(Equal(localhost.Ip6))
	})

	It("resolves dns name", func() {
		pinger, _ := NewPinger(context.TODO(), localhost.DnsName)
		Expect(pinger.IPAddr().IP).To(Equal(localhost.Ip4))
	})

	Context("invalid host", func() {
		expectedError := func(name string) (*net.DNSError){
			return &net.DNSError{Err: "no such host", Name: name}
		}
		Describe("invalid hostname", func() {
			It("returns lookup error", func() {
				host := "lonelyhost"
				_, err := NewPinger(context.TODO(), host)
				Expect(err).To(Equal(expectedError(host)))
			})
		})

		Describe("invalid ip address", func() {
			It("returns invalid ip error", func() {
				host := ":::1"
				_, err := NewPinger(context.TODO(), host)
				Expect(err).To(Equal(expectedError(host)))
			})
		})
	})
})

var _ = Describe("NewPingerWithNetwork", func() {

	It("returns pointer to a new Pinger", func() {
		pinger, _ := NewPingerWithNetwork(context.TODO(), localhost.DnsName, "ip")
		Expect(pinger).ToNot(BeNil())
		Expect(pinger).To(BeAssignableToTypeOf(&Pinger{}))
	})

	Describe("enforce ip4", func() {
		It("resolves dns to ip4 address", func() {
			pinger, _ := NewPingerWithNetwork(context.TODO(), localhost.DnsName, "ip4")
			Expect(pinger.IPAddr().IP).To(Equal(localhost.Ip4))
		})
	})

	Describe("enforce ip6", func() {
		It("resolves dns to ip6 address", func() {
			pinger, _ := NewPingerWithNetwork(context.TODO(), localhost.DnsName, "ip6")
			Expect(pinger.IPAddr().IP).To(Equal(localhost.Ip6))
		})
	})

})

var _ = Describe("Pinger", func() {
	Describe("Run()", func() {
		PIt("test execution")
		PIt("test multithreaded execution/ package id matching")
		PIt("test stopping deadlock")
		PIt("test timer leakage")
	})

	Describe("Statistics()", func() {
		PIt("test calculation")
	})
})
