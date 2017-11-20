package cli

type BeforeFunc func() error

type AfterFunc func() error
