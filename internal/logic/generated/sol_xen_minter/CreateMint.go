// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package sol_xen_minter

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CreateMint is the `createMint` instruction.
type CreateMint struct {
	Metadata *InitTokenParams

	// [0] = [WRITE, SIGNER] admin
	//
	// [1] = [WRITE] mintAccount
	//
	// [2] = [] tokenProgram
	//
	// [3] = [] systemProgram
	//
	// [4] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCreateMintInstructionBuilder creates a new `CreateMint` instruction builder.
func NewCreateMintInstructionBuilder() *CreateMint {
	nd := &CreateMint{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetMetadata sets the "metadata" parameter.
func (inst *CreateMint) SetMetadata(metadata InitTokenParams) *CreateMint {
	inst.Metadata = &metadata
	return inst
}

// SetAdminAccount sets the "admin" account.
func (inst *CreateMint) SetAdminAccount(admin ag_solanago.PublicKey) *CreateMint {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(admin).WRITE().SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *CreateMint) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMintAccountAccount sets the "mintAccount" account.
func (inst *CreateMint) SetMintAccountAccount(mintAccount ag_solanago.PublicKey) *CreateMint {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(mintAccount).WRITE()
	return inst
}

// GetMintAccountAccount gets the "mintAccount" account.
func (inst *CreateMint) GetMintAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *CreateMint) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *CreateMint {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *CreateMint) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *CreateMint) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *CreateMint {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *CreateMint) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetRentAccount sets the "rent" account.
func (inst *CreateMint) SetRentAccount(rent ag_solanago.PublicKey) *CreateMint {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *CreateMint) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst CreateMint) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CreateMint,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CreateMint) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CreateMint) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Metadata == nil {
			return errors.New("Metadata parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.MintAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *CreateMint) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CreateMint")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Metadata", *inst.Metadata))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        admin", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("         mint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta(" tokenProgram", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("         rent", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj CreateMint) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Metadata` param:
	err = encoder.Encode(obj.Metadata)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CreateMint) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Metadata`:
	err = decoder.Decode(&obj.Metadata)
	if err != nil {
		return err
	}
	return nil
}

// NewCreateMintInstruction declares a new CreateMint instruction with the provided parameters and accounts.
func NewCreateMintInstruction(
	// Parameters:
	metadata InitTokenParams,
	// Accounts:
	admin ag_solanago.PublicKey,
	mintAccount ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *CreateMint {
	return NewCreateMintInstructionBuilder().
		SetMetadata(metadata).
		SetAdminAccount(admin).
		SetMintAccountAccount(mintAccount).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}
