# Chapter 10. Tokens

- Tokens are commonly used to refer to privately issued special-purpose coin-like items of insignificant intrinsic value, such as transportation tokens, laundry tokens, and arcade game tokens.
- In blockchain, tokens are blockchain-based abstractions that can be owned and that represent assets, currency, or access rights

## How Tokens Are Used

| Application | Description                                                                                                                                                                                             |
| ----------: | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
|    Currency | A token can serve as a form of currency, with a value determined through private trade. **Most obviously used**                                                                                         |
|    Resource | A token can represent a resource earned or produced in a sharing economy or resource-sharing environment; for example, a storage or CPU token representing resources that can be shared over a network. |
|       Asset | A token can represent ownership of an intrinsic or extrinsic, tangible or intangi‐ ble asset; for example, gold, real estate, a car, oil, energy, MMOG items, etc.                                      |
|      Access | A token can represent access rights and grant access to a digital or physical prop‐ erty, such as a discussion forum, an exclusive website, a hotel room, or a rental car.                              |
|      Equity | A token can represent shareholder equity in a digital organization (e.g., a DAO) or legal entity (e.g., a corporation).                                                                                 |
|      Voting | A token can represent voting rights in a digital or legal system.                                                                                                                                       |
| Collectible | A token can represent a digital collectible (e.g., CryptoPunks) or physical collec‐ tible (e.g., a painting).                                                                                           |
|    Identity | A token can represent a digital identity (e.g., avatar) or legal identity (e.g., national ID).                                                                                                          |
| Attestation | A token can represent a certification or attestation of fact by some authority or by a decentralized reputation system (e.g., marriage record, birth certificate, college degree).                      |
|     Utility | A token can be used to access or pay for a service.                                                                                                                                                     |

## Tokens and Fungibility

> WIKI: In economics, fungibility is the property of a good or a commodity whose individual units are essentially interchangeable

- Tokens are **fungible** when we can substitute any single unit of the token for another without any difference in its value or function
- If a token's historical provenance can be tracked, then it is not entirely fungible
  - The ability to track provenance can lead to blacklisting and whitelisting, reducing or eliminating fungibility
- Non-fungible tokens are tokens that each represent a unique tangible or intangible item and therefore are not interchangeable
  - Each non-fungible token is associated with a unique identifier, such as a serial number

> Note that "fungible" is often used to mean "directly exchangeable for money"

## Counterparty Risk

- **WHAT**: The risk that the other party in a transaction will fail to meet their obligations, usually involving more than 2 parties
- **WHEN**: When an asset is traded indirectly through the exchange of a token of ownership, there is additional counterparty risk from the custodian of the asset
- In the world of digital tokens representing assets, it is important to understand
  - Who holds the asset that is represented by the token
  - What rules apply to that underlying asset

## Tokens and Intrinsicality

- Some tokens represent digital items that are intrinsic to the blockchain. Those digital assets are governed by consensus rules, resisting against additional counterparty risk
- The ownership of extrinsic items off-chain, is governed by law, custom, and policy, separate from the consensus rules that govern the token, suffering from additional counterparty risk
  - Examples: real estate, corporate voting shares, trademarks, and gold bars
- One of the most important ramifications of blockchain-based tokens is the ability to convert extrinsic assets into intrinsic assets and thereby remove counterparty risk
  - A good example is moving from equity in a corporation (extrinsic) to an equity or vot‐ ing token in a DAO or similar (intrinsic) organization

## Using Tokens: Utility or Equity

- **Utility tokens**
  - **WHAT**: The use of the token is required to gain access to a service, application, or resource
  - Examples: Tokens that represent resources such as shared storage, or access to services such as social media networks
- **Equity tokens**
  - **WHAT**: Tokens representing shares in the control or ownership of something
  - Examples:
    - Nonvoting shares for dis‐ tribution of dividends and profits
    - Voting shares in DAO

### It's a Duck!

- Many startups face a difficult problem
  - Tokens are a great fundraising mechanism
  - But offering securities (equity) to the public is a regulated activity in most jurisdictions
- Disguising equity tokens as utility tokens is watched closely by regulators

### Utility Tokens: Who Needs Them?

- For a startup, each innovation represents a **risk** and a **market filter**
- Adding a utility token to that innovation and requiring users to adopt tokens in order to use the service compounds the risk and increases the barriers to adoption
- As a **filter**
  - Innovation limits adoption to the subset of the market that can become early adopters of this innovation
  - Adding a second filter compounds that effect, further limiting the addressable market
- **Risk** from
  - The underlying platform (Ethereum)
  - Broader economy (exchanges, liquidity)
  - Regulatory environment (equity/commodity regulators)
  - Technology (smart contracts, token standards)
- Pro: By adopting tokens they are also inheriting the market enthusiasm, early adopters, technology, innovation, and liquidity of the entire token economy
- Limited liquidity, limited applicability, and high conversion costs reduce the value of tokens until they are only of "token" value
- The switching costs of a digital token are orders of magnitude lower than for a physical token without a market, but they are not zero
- Advices
  - Adopt a token because your application cannot work without a token
  - Adopt it because the token lifts a fundamental market barrier or solves an access problem
  - Don't introduce a utility token because it is the only way you can raise money fast and you need to pretend it's not a public securities offering

## Tokens on Ethereum

- The ether balance of Ethereum accounts is handled at the protocol level
- The token balance of Ethereum accounts is handled at the smart contract level

### The ERC20 Token Standard

- **WHEN**: Introduced in November 2015 by Fabian Vogelsteller
- The ERC20 request for comments eventually became [EIP-20](http://eips.ethereum.org/EIPS/eip-20)
- ERC20 is a standard for fungible tokens, meaning that different units of an ERC20 token are interchangeable and have no unique properties
- The ERC20 standard defines a common interface for contracts implementing a token, such that any compatible token can be accessed and used in the same way

#### Specification

```solidity
contract ERC20 {
  // Returns the total token supply
  function totalSupply() constant returns (uint theTotalSupply);

  // Returns the account balance of another account with address _owner
  function balanceOf(address _owner) constant returns (uint balance);

  // Transfers _value amount of tokens to address _to, and MUST fire the Transfer event.
  // The function SHOULD throw if the message caller’s account balance does not have enough tokens to spend
  function transfer(address _to, uint _value) returns (bool success);

  // Transfers _value amount of tokens from address _from to address _to, and MUST fire the Transfer event.
  function transferFrom(address _from, address _to, uint _value) returns (bool success);

  // Allows _spender to withdraw from your account multiple times, up to the _value amount. If this function
  // is called again it overwrites the current allowance with _value.
  function approve(address _spender, uint _value) returns (bool success);

  // Returns the amount which _spender is still allowed to withdraw from _owner.
  function allowance(address _owner, address _spender) constant returns (uint remaining);

  // MUST trigger when tokens are transferred, including zero value transfers
  event Transfer(address indexed _from, address indexed _to, uint _value);

  // MUST trigger on any successful call to approve()
  event Approval(address indexed _owner, address indexed _spender, uint _value);

  // 3 more OPTIONAL functions

  // Returns the name of the token - e.g. "MyToken"
  function name() public view returns (string)

  // Returns the symbol of the token. E.g. "HIX"
  function symbol() public view returns (string)

  // Returns the number of decimals the token uses - e.g. 8, means to divide the token amount by 100000000
  // to get its user representation
  function decimals() public view returns (uint8)
}
```

#### Data Structures

- `balances`
  - Type: `mapping(address => uint256)`
  - **FOR**: Keep track of who owns the tokens
- `allowances`
  - Type: `mapping (address => mapping (address => uint256)) public`
  - **WHAT**: With the primary key being the address of the token owner, mapping to a spender address and an allowance amount

#### ERC20 workflows: "transfer" and "approve & transferFrom"

- `transfer`
  - A single-tx workflow
  - **WHEN**: Exchanging tokens between wallets
- `approve` followed by `transferFrom`

  - A two-tx workflow
  - **WHEN**: Allow a token owner to delegate their control to another address depicted as follows
    ![The two-step approve & transferFrom work ow of ERC20 tokens](images/approve_transferFrom_workflow.png)

    > It is most often used to delegate control to a contract for distribution of tokens, but it can also be used by exchanges

  - **USE CASE**: ICO

#### Implementation

- [Consensys EIP20](https://github.com/ConsenSys/Tokens/tree/master/contracts/eip20)
- [OpenZeppelin StandardToken](https://github.com/OpenZeppelin/openzeppelin-solidity/blob/v1.12.0/contracts/token/ERC20/StandardToken.sol)

### Launching Our Own ERC20 Token

1. Make the directory

   ```bash
   mkdir METoken
   ```

2. Install dependencies

   ```bash
   yarn add openzeppelin-solidity@2.2.0
   ```

3. Compile the contract by running the [compile.sh](examples/METoken/compile.sh) script
   > The bytecodes for each component would be produced in the `contracts/build` folder
4. Deploy the contract with script [deploy.go](examples/METoken/deploy.go)
5. Once the contract is settled on the blockchain, check it with [deploy_test.go](examples/METoken/deploy_test.go)

#### Interacting with METoken

1. Transfer some tokens between accounts as [transfer.go](examples/METoken/transfer.go)
2. Once the transfering is settled, check it by script [transfer_test.go](examples/METoken/transfer_test.go)

#### Sending ERC20 tokens to contract addresses

- TODO: demo code

> The Faucet contractdoesn't have a function for withdrawing MET, or any other ERC20 token. If we use `withdraw` it will try to send ether, but since Faucet doesn't have a balance of ether yet, it will fail.

- One of the ways that users of ERC20 tokens can inadvertently lose their tokens in a transfer, is when they attempt to transfer to an exchange or another service

#### Demonstrating the "approve & transferFrom" workflow

- TODO: demo code

### Issues with ERC20 Tokens

- One of the less obvious issues is that they expose subtle differences between tokens and ether itself
  - Token transfers occur within the specific token contract state and have the token contract as their destination, not the recipient's address
  - In a token transfer, NO TRANSACTION IS ACTUALLY SENT TO THE RECIPIENT OF THE TOKEN
  - Even a wallet that has support for ERC20 tokens does not become aware of a token balance unless the user explicitly adds a specific token contract to "watch."
- Tracking all balances in all possible ERC20 token contracts would suffer from problems like email spam
- Ether is sent with the `send` function and accepted by any `payable` function in a contract or any externally owned address. Tokens are sent using `transfer` or `approve` & `transferFrom` functions that exist only in the ERC20 contract, and do not (at least in ERC20) trigger any payable functions in a recipient contract
- Sending tokens require ethers, causing some strange UX

### ERC223: A Proposed Token Contract Interface Standard

- **WHY**: Solve the problem of inadvertent transfer of tokens to a contract (that may or may not support tokens) by detecting whether the destination address is a contract or not
- **HOW**: Contracts designed to accept tokens must implement a function named `tokenFallback`
  - If `tokenFallback` is missing, the transfer fails
- Contract detection snippet

  ```solidity
  function isContract(address _addr) private view returns (bool is_contract) {
    uint length;
    assembly {
      // retrieve the size of the code on target address; this needs assembly
      length := extcodesize(_addr)
    }
    return (length>0);
  }
  ```

- Full spec goes as

  ```solidity
  interface ERC223Token {
    uint public totalSupply;
    function balanceOf(address who) public view returns (uint);

    function name() public view returns (string _name);
    function symbol() public view returns (string _symbol);
    function decimals() public view returns (uint8 _decimals);
    function totalSupply() public view returns (uint256 _supply);

    function transfer(address to, uint value) public returns (bool ok);
    function transfer(address to, uint value, bytes data) public returns (bool ok);
    function transfer(address to, uint value, bytes data, string custom_fallback) public returns (bool ok);

    event Transfer(address indexed from, address indexed to, uint value, bytes indexed data);
  }
  ```

- ERC223 is not widely implemented, due to

### ERC777: A Proposed Token Contract Interface Standard

- **WHY**

  - To offer an ERC20-compatible interface
  - To transfer tokens using a send function, similar to ether transfers
  - To be compatible with ERC820 for token contract registration
  - To allow contracts and addresses to control which tokens they send through a tokensToSend function that is called prior to sending
  - To enable contracts and addresses to be notified of the tokens’ receipt by calling a tokensReceived function in the recipient, and to reduce the probability of tokens being locked into contracts by requiring contracts to provide a tokensReceived function
  - To allow existing contracts to use proxy contracts for the tokensToSend and tokensReceived functions
  - To operate in the same way whether sending to a contract or an EOA
  - To provide specific events for the minting and burning of tokens
  - To enable operators (trusted third parties, intended to be verified contracts) to move tokens on behalf of a token holder
  - To provide metadata on token transfer transactions in userData and operator Data fields

- Specification as [ERC777](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-777.md)
- Some of the debate on ERC777 is about the complexity of adopting two big changes at once: a new token standard and a registry standard

### ERC721: Non-fungible Token (Deed) Standard

- **WHAT**: A standard for non-fungible tokens, also known as deeds, spcified in [EIP-721](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-721.md)
  > From the Oxford Dictionary: deed: A legal document that is signed and delivered, especially one regarding the ownership of property or legal rights.
- Non-fungible tokens track ownership of a unique thing
  - The ERC721 requires only that the tracked items can be uniquely identified, by means of a mapping from 256-bit deed ID to item ID

## Using Token Standards

### What Are Token Standards? What Is Their Purpose?

- **WHAT**: Token standards are the minimum specifications for an implementation
- **WHY**: To encourage interoperability between contracts
- **CAVEAT**: The standards are meant to be **descriptive**, rather than prescriptive
  - The implementation is up to the developers.

### Should You Use These Standards?

- **A dilemma**: use the existing standards or innovate beyond the restrictions they impose?
- **WHY STANDARDS**: Interoperability and broad adoption
- **Not Invented Here** syndrome
  - **WHAT**: The tendency to forge your own path and ignore existing standards
  - It's antithetical to open source culture
  - However, progress and innovation depend on departing from tradition sometimes

### Security by Maturity

- Existing implementations are "battle-tested."
  - While it is impossible to prove that they are secure, many of them underpin millions of dollars' worth of tokens.
  - They have been attacked, repeatedly and vigorously
  - It is much safer to use a well-tested, widely used implementation
- Extending existing implementations would introduce extra complexity, further expanding the attack surface

## Extensions to Token Interface Standards

- **Owner control**: The ability to give specific addresses, or sets of addresses (i.e., multisignature schemes), special capabilities, such as blacklisting, whitelisting, minting, recov‐ ery, etc.
- **Burning**: The ability to deliberately destroy (“burn”) tokens by transferring them to an unspendable address or by erasing a balance and reducing the supply
- **Minting**: The ability to add to the total supply of tokens, at a predictable rate or by “fiat” of the creator of the token.
- **Crowdfunding**: The ability to offer tokens for sale, for example through an auction, market sale, reverse auction, etc.
- **Caps**: The ability to set predefined and immutable limits on the total supply (the oppo‐ site of the “minting” feature).
- **Recovery backdoors**: Functions to recover funds, reverse transfers, or dismantle the token that can be activated by a designated address or set of addresses.
- **Whitelisting**: The ability to restrict actions (such as token transfers) to specific addresses. Most commonly used to offer tokens to “accredited investors” after vetting by the rules of different jurisdictions. There is usually a mechanism for updating the white‐ list.
- **Blacklisting**: The ability to restrict token transfers by disallowing specific addresses. There is usually a function for updating the blacklist.

> The decision to extend a token standard with additional functionality represents a trade-off between innovation/risk and interoperability/security.

## Tokens and ICOs

- Many of the tokens on offer in Ethereum today are barely disguised scams, pyramid schemes, and money grabs

## Conclusions
